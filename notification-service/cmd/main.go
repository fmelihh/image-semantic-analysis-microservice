package main

import (
	"fmt"
	"notification-service/config"
	"notification-service/service"
	"notification-service/types"
	"notification-service/utils"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	kafkaConsumer := service.NewKafkaConsumerService()
	smtpConfigurations := types.SmtpConfigurations{
		Host:        config.Envs.SMTPHost,
		Port:        config.Envs.SMTPPort,
		Login:       config.Envs.SMTPLogin,
		AccessToken: config.Envs.SMTPToken,
	}
	notificationService := service.NewNotificationService(smtpConfigurations)

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
	errorMsgCount := 0
	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				fmt.Printf("Received Message Count: %d: | Topic (%s) | Message(%s)\n", msgCount, string(msg.Topic), string(msg.Value))
				convertedMessage := utils.ConvertBytesToMap(msg.Value)
				notifiedEmail, err := notificationService.Notify(convertedMessage)
				if err != nil {
					errorMsgCount++
					fmt.Printf("An error occurred. Details: %v \n", err.Error())
				} else {
					msgCount++
					fmt.Printf("Message Notified to %s \n", notifiedEmail)
				}

				fmt.Printf("Total received messages %d, Total failed messages %d \n", msgCount, errorMsgCount)

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
