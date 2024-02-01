package greetings

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"

)
const TOPIC = "string"

func createConsumer() {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092",
		"group.id": "cluster", 
		"api.version.request": true,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		fmt.Println(err)
	}

	err = consumer.Subscribe(TOPIC, nil)
	if err != nil {
		fmt.Printf("Erreur lors de l'abonnement aux sujets: %v\n", err)
		return
	}

	defer consumer.Close()

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message reçu: %s\n", string(msg.Value))
		} else {
			fmt.Printf("Erreur lors de la lecture du message: %v\n", err)
		}
	}	
}	