package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	kafkaProduce "github.com/rhadamez/simulator-aluno/application/kafka"
	affe "github.com/rhadamez/simulator-aluno/infra/kafka"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	msgChan := make(chan *ckafka.Message)
	consumer := affe.NewKafkaConsumer(msgChan)
	go consumer.Consume()

	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafkaProduce.Produce(msg)
	}
}
