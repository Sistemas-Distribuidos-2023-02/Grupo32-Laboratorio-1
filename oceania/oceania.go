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
	central = flag.String("addr_central", "dist125.inf.santiago.usm.cl:50051", "the address to connect to")
	oceania = flag.Int("oceania_port", 50055, "The server port")
	usuarios int
	interesados int
	min int
	max int 
)

type server struct {
	pb.UnimplementedServersServiceServer
}

func (s *server) MandarLlaves(ctx context.Context, in *pb.Llaves) (*pb.Confirmar, error) {
	go rabbit()
	return &pb.Confirmar{Flag: "1"}, nil
}
func (s server) MandarNoAccedidos(ctx context.Context, in *pb.Llaves) (*pb.Confirmar, error) {
	runes := []rune(in.GetNumero())
    faltantes := int(runes[0]) //conversión mágica	
	
    usuarios = int(math.Max(0, float64(usuarios - (interesados - faltantes)))) 
	
	log.Printf("Usuarios que lograron inscribirse: %d", interesados - faltantes)
	log.Printf("Usuarios restantes en espera de cupo : %d", usuarios)
	min = int(float64(usuarios)*0.4)
	max = int(float64(usuarios)*0.6)

    return &pb.Confirmar{Flag: "1"}, nil
}


func main() {
	flag.Parse()

	file, _ := os.Open("oceania/parametros_de_inicio.txt")
	scanner := bufio.NewScanner(file)
	//scanner.Split(bufio.ScanWords)
	scanner.Scan()
	aux, _ := strconv.Atoi(scanner.Text())
	log.Printf( "Hay ", strconv.Itoa(aux) ," personas totales interesadas en inscribir la beta" )
	min = int(float64(aux)*0.4)
	max = int(float64(aux)*0.6)
	usuarios = aux

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *oceania))
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
	conn, err := amqp.Dial("amqp://guest:guest@dist126.inf.santiago.usm.cl:5672/")
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
	interesados = llaves
	s := [2]string{"3", strconv.Itoa(llaves)}

	envio := strings.Join(s[0:], " ")
	log.Printf("Se solicita a la central inscribir ", llaves, " cupos")

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
