package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	proto := getEnv("RABBITMQ_PROTO", "amqps")
	host := getEnv("RABBITMQ_HOST", "localhost")
	port := getEnv("RABBITMQ_PORT", "5672")
	user := getEnv("RABBITMQ_USER", "user")
	password := getEnv("RABBITMQ_PASSWORD", "password")
	amqpURL := fmt.Sprintf("%s://%s:%s@%s:%s", proto, user, password, host, port)

	// Connect to RabbitMQ
	conn, err := amqp091.Dial(amqpURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()
	log.Printf("Connected to RabbitMQ %s://%s:%s", proto, host, port)

	// Open a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()
	log.Println("Channel opened")

	// Declare a queue (optional, but common)
	queueName := "test_queue"
	_, err = ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}
	log.Printf("Queue %s declared", queueName)

	ch.QueueDelete("test_queue", false, false, false)

	log.Printf("Queue %s deleted", queueName)
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
