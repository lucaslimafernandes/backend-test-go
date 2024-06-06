package main

import (
	"backendtest-go/models"
	"backendtest-go/services"
	"fmt"
	"log"
)

func init() {
	models.LoadEnvs()
	models.RConn()
}

func main() {

	go notifify()
	go delete()

	// run forever
	select {}

}

func notifify() {

	msgs, err := models.RabbitMQChannel.Consume(
		models.NotifyQueue.Name, // queue
		"",                      // consumer
		true,                    // auto-ack
		false,                   // exclusive
		false,                   // no-local
		false,                   // no-wait
		nil,                     // args
	)
	if err != nil {
		log.Println("Error consuming notify queue:", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			services.SendEmail(d.Body)
		}
	}()

	fmt.Println("Waiting for notify messages")
	<-forever

}

func delete() {

	msgs, err := models.RabbitMQChannel.Consume(
		models.UnsafeQueue.Name, // queue
		"",                      // consumer
		true,                    // auto-ack
		false,                   // exclusive
		false,                   // no-local
		false,                   // no-wait
		nil,                     // args
	)
	if err != nil {
		log.Println("Error consuming delete queue:", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			services.DeleteFile(d.Body)
		}
	}()

	fmt.Println("Waiting for delete messages")
	<-forever

}
