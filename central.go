package main

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"log"
	"math/rand"
    "time"
	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	central = flag.Int("central_port", 50051, "The central port")
	asia = flag.String("addr_asia", "localhost:50052", "the address to connect to")
)

type server struct {
	pb.UnimplementedChatServiceServer
}


func (s *server) MandarLlaves(ctx context.Context, in *pb.Llaves) (*pb.Confirmar, error) {
	log.Printf("Mandado: %v", in.GetFlag())
	return &pb.Confirmar{flag: "Numero de llaves recibidas: " + in.GetNumero()}, nil
}

func main(){
	file, _ := os.Open("parametros_de_inicio.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	rango := strings.Split(scanner.Text(), "-")
	min, _:= strconv.Atoi(rango[0])
	max, _ := strconv.Atoi(rango[1])
	scanner.Scan()
	iteraciones, _ := strconv.Atoi(scanner.Text())
	file.Close()

	/*
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

	q, err := ch.QueueDeclare(
		"asia", 
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

	log.Println(q)

	*/

	// Asia
	conn_asia, err := grpc.Dial(*asia, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn_asia.Close()
	c := pb.NewServersServiceConfirmar(conn_asia)

	// America
	// Europa
	// Oceania


	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	

	if iteraciones != -1 {
		for i := 0; i < iteraciones; i++ {
			rand.Seed(time.Now().UnixNano())
			llaves := rand.Intn(max - min + 1) + min

			// Asia
			r, err := c.MandarLlaves(ctx, &pb.Llaves{Numero: *llaves})  // x4
			if err != nil {
				log.Fatalf("could not send: %v", err)
			}
			log.Printf("Sending: %s", r.GetBody())
					
			// America
			// Europa
			// Oceania

			// Rabbit




		}
	} // else

	

	if err != nil {
		log.Println(err)
		panic(err)
	}

	log.Println("listo")


}