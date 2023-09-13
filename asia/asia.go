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

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/LaTortugaR/testing_something/protos"
)

const (
	defaultUsers = "0"
)

var (
	central = flag.String("addr_central", "localhost:50051", "the address to connect to")
	asia = flag.Int("asia_port", 50052, "The server port")
	usuarios = flag.String("usuarios", defaultUsers, "Cantidad de usuarios")
)

func main() {
	flag.Parse()

	file, _ := os.Open("parametros_de_inicio.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	usuarios := strconv.Atoi(scanner.Text())
	min := usuarios*0.4
	max := usuarios*0.6


	/// Pausa

	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewChatServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.Request{Cliente: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetBody())
}