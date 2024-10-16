package main

import (
	"fmt"
	"notification-service/service"
	"notification-service/utils"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	kafkaConsumer := service.NewKafkaConsumerService()
	notificationService := service.NewNotificationService()

	conn, err := kafkaConsumer.ConnectConsumer([]string{"localhost:29092"})
	if err != nil {
		panic(err)
	}
	consumer, err := kafkaConsumer.SubscribeTopic(conn, "notification")

	if err != nil {
		panic(err)
	}

	fmt.Println("Notification consumer started.")

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	msgCount := 0
	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				msgCount++
				fmt.Printf("Received Message Count: %d: | Topic (%s) | Message(%s)\n", msgCount, string(msg.Topic), string(msg.Value))

				convertedMessage := utils.ConvertBytesToMap(msg.Value)
				notificationService.Notify(convertedMessage["email"])

			case <-sigchan:
				fmt.Println("Interrupted detected.")
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	fmt.Println("Processed", msgCount, "messages")
	if err := consumer.Close(); err != nil {
		panic(err)
	}
}
