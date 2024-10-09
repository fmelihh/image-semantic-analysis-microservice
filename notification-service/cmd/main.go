package main

import (
	"fmt"
	"notification-service/service"
)

func main() {
	kafkaConsumer := service.NewKafkaConsumerService()
	conn, err := kafkaConsumer.ConnectConsumer([]string{"localhost:29092"})
	if err != nil {
		panic(err)
	}
	consumer, err := kafkaConsumer.SubscribeTopic(conn, "notification")

	if err != nil {
		panic(err)
	}

	fmt.Println("Consumer started.")

}
