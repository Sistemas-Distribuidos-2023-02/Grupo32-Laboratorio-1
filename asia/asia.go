package main

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"context"
	"flag"
	"log"
	"time"
	"math/rand"
	"net"
	"fmt"
	"math"

	"google.golang.org/grpc"
	pb "github.com/LaTortugaR/ProtosLab1/protos"
	amqp "github.com/rabbitmq/amqp091-go"
)


var (
	central = flag.String("addr_central", "localhost:50051", "the address to connect to")
	asia = flag.Int("asia_port", 50052, "The server port")
	usuarios int
	interesados int
	min int
	max int 
)

type server struct {
	pb.UnimplementedServersServiceServer
}

func (s *server) MandarLlaves(ctx context.Context, in *pb.Llaves) (*pb.Confirmar, error) {
	log.Printf("Recibido: %d", in.GetNumero())
	go rabbit()
	return &pb.Confirmar{Flag: "1"}, nil
}
func (s server) MandarNoAccedidos(ctx context.Context, in *pb.Llaves) (*pb.Confirmar, error) {
    log.Printf("Recibido: %d", in.GetNumero())
    usuariosValor := usuarios
    interesadosValor, err := strconv.Atoi(in.GetNumero()) 
    if err != nil{
        return nil, err
    }
    usuariosValor = int(math.Max(0, float64(usuariosValor - (interesadosValor - usuariosValor))))
    usuarios = usuariosValor

    return &pb.Confirmar{Flag: "1"}, nil
}


func main() {
	flag.Parse()

	file, _ := os.Open("asia/parametros_de_inicio.txt")
	scanner := bufio.NewScanner(file)
	//scanner.Split(bufio.ScanWords)
	scanner.Scan()
	usuarios, _ := strconv.Atoi(scanner.Text())
	log.Println(usuarios)
	min = int(float64(usuarios)*0.4)
	max = int(float64(usuarios)*0.6)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *asia))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterServersServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	/// GRPC escucha 1º vez
}

func rabbit() {

	//Rabbit
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Println(err)
		panic(err)
	}

	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		log.Println(err)
		panic(err)
	}

	defer ch.Close()

	_, err = ch.QueueDeclare(
		"cola", 
		false, 
		false,   
		false,   
		false,   
		nil,     
	)

	if err != nil {
		log.Println(err)
		panic(err)
	}


	rand.Seed(time.Now().UnixNano())
	llaves := rand.Intn(max - min + 1) + min
	s := [2]string{"AS", string(llaves)}

	envio := strings.Join(s[1:], " ")
	log.Println("llaves: ",llaves)

	err = ch.Publish(
		"",
		"cola",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(envio),
		},
	)

}
