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
	"math"

    "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/LaTortugaR/ProtosLab1/protos"
	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	central = flag.Int("central_port", 50051, "The central port")
	asia = flag.String("addr_asia", "localhost:50052", "the address to connect to")
	europa = flag.String("addr_europa", "localhost:50053", "the address to connect to")
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

	llaves := 0
	servers := [2]string{"asia", "europa"}
	

	// Asia
	conn_asia, err := grpc.Dial(*asia, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn_asia.Close()
	c_asia := pb.NewServersServiceClient(conn_asia)

	// Europa
	conn_europa, err := grpc.Dial(*europa, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn_asia.Close()
	c_europa := pb.NewServersServiceClient(conn_europa)

	// America
	
	// Oceania



	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	
	canal := make(chan int)
	
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
	
	sum := 0

	
	msgs, _ := ch.Consume(
		"cola",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	confirmar := [4]int{1,1,1,1}
	
	go func() {
		for d := range msgs {
			region, _ := strconv.Atoi(string(d.Body[0]))  
			log.Println("Llaves pedidas por", servers[region],": ", string(d.Body[1:]))
			stringg := string(d.Body)
			stringg = stringg[2:] 

			llaves_pedidas := 0
			if len(stringg) != 0 {
				llaves_pedidas, _ = strconv.Atoi(stringg)
				confirmar[region], _ = strconv.Atoi(stringg)
			} else{
				llaves_pedidas = 0

			}

			dif := llaves_pedidas - llaves
			no_accedidos := int(math.Max(math.Abs(float64(dif)), 0.0))   // cuantos no alcanzaron
			llaves = int( math.Min(math.Abs(float64(-dif)), 0.0) )  //cuantos quedan disponibles

			//MANDAR LLAVES CON GRPC
			if region == 0 {
				_, err := c_asia.MandarNoAccedidos(ctx, &pb.Llaves{Numero: string(no_accedidos)})  // x4
				if err != nil {
					log.Fatalf("could not send: %v", err)
				}
				
				log.Println("Llaves q no accedio asia: ",no_accedidos)
			} else if region == 1 {
				_, err := c_europa.MandarNoAccedidos(ctx, &pb.Llaves{Numero: string(no_accedidos)})  // x4
				if err != nil {
					log.Fatalf("could not send: %v", err)
				}
				log.Println("Llaves q no accedio europa: ",no_accedidos)
			}

			j := 0
			for m := 0; m < len(confirmar); m++ {
				if confirmar[m] == 0{j++}
			}
			if j == 4{iteraciones = 0}
			canal <- 1

		}	
	}()

	if iteraciones != -1 {
		for i := 0; i < iteraciones; i++ {
			rand.Seed(time.Now().UnixNano())
			llaves = rand.Intn(max - min + 1) + min
			log.Println("Llaves creadas en la iteracion " + strconv.Itoa(i+1) + ": " + strconv.Itoa(llaves))

			// Asia
			r, err := c_asia.MandarLlaves(ctx, &pb.Llaves{Numero: string(llaves)})  // x4
			if err != nil {
				log.Fatalf("could not send: %v", err)
			}
			log.Printf("Sending: %s", r.GetFlag())
					
			// Europa
			r, err = c_europa.MandarLlaves(ctx, &pb.Llaves{Numero: string(llaves)})  // x4
			if err != nil {
				log.Fatalf("could not send: %v", err)
			}
			log.Printf("Sending: %s", r.GetFlag())


			// America
			// Oceania

			// Rabbit
			
			
	
			
			for ; sum < 2; {
				select {
				case  <- canal:
					sum++
				}
			}

			sum = 0
		

		}
	}else{
		for ; iteraciones == 0; {
			rand.Seed(time.Now().UnixNano())
			llaves = rand.Intn(max - min + 1) + min

			// Asia
			r, err := c_asia.MandarLlaves(ctx, &pb.Llaves{Numero: string(llaves)})  // x4
			if err != nil {
				log.Fatalf("could not send: %v", err)
			}
			log.Printf("Sending: %s", r.GetFlag())
					
			// Europa
			r, err = c_europa.MandarLlaves(ctx, &pb.Llaves{Numero: string(llaves)})  // x4
			if err != nil {
				log.Fatalf("could not send: %v", err)
			}
			log.Printf("Sending: %s", r.GetFlag())


			// America
			// Oceania

			// Rabbit
			
			
	
			
			for ; sum < 2; {
				select {
				case  <- canal:
					sum++
				}
			}

			sum = 0
		

		}
	}
		
	

	if err != nil {
		log.Println(err)
		panic(err)
	}

	log.Println("listo")


}
