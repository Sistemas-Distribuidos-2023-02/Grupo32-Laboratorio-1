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

	if iteraciones != -1{
		for i := 0; i < iteraciones; i++ {
			rand.Seed(time.Now().UnixNano())
			llaves := rand.Intn(max - min + 1) + min
			log.Println(llaves)
			err = ch.Publish(
				"",
				"asia",
				false,
				false,
				amqp.Publishing{
					ContentType: "text/plain",
					Body: []byte(string(llaves)),
				},
			)
					
		}
	}

	

	if err != nil {
		log.Println(err)
		panic(err)
	}

	log.Println("listo")


}