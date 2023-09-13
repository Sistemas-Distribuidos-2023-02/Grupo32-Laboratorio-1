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

    "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/LaTortugaR/ProtosLab1/protos"
	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	central = flag.Int("central_port", 50051, "The central port")
	asia = flag.String("addr_asia", "localhost:50052", "the address to connect to")
)


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

	// Asia
	conn_asia, err := grpc.Dial(*asia, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn_asia.Close()
	c := pb.NewServersServiceClient(conn_asia)

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
			r, err := c.MandarLlaves(ctx, &pb.Llaves{Numero: string(llaves)})  // x4
			if err != nil {
				log.Fatalf("could not send: %v", err)
			}
			log.Printf("Sending: %s", r.GetFlag())
					
			// America
			// Europa
			// Oceania

			// Rabbit
			
			ch := rabbit()

			for j := 0; j < 4; {
				msgs, _ := ch.Consume(
					"cola",
					"",
					true,
					false,
					false,
					false,
					nil,
				)
				
				go func() {
					for d := range msgs {
						log.Println(d.Body)
						j++;
					}
				}()

				//MANDAR LLAVES CON GRPC
			}



		}
	} // else

	

	if err != nil {
		log.Println(err)
		panic(err)
	}

	log.Println("listo")


}

func rabbit() (*amqp.Channel){
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
	
	return ch
}
