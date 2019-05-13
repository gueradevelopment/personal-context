package services

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"log"
)

type RabbitService struct {
	connection *amqp.Connection
}

func RabbitServiceInit() *RabbitService {
	service := new(RabbitService)
	connection, err := amqp.Dial("amqp://guera:pass@34.73.163.109")
	if err != nil {
		log.Fatalln("Rabbit Connection Error")
		return nil
	}
	service.connection = connection
	return service
}

func (service *RabbitService) SendAndReceive(message string, routingKey string, c chan string) {
	defer close(c)
	channel, err := service.connection.Channel()
	if err != nil {
		return
	}
	defer channel.Close()
	channel.ExchangeDeclare(
		"data-provision", "topic", true, false, false, false, nil,
	)

	correlationId := uuid.New().String()
	queueName := uuid.New().String()
	q, err := channel.QueueDeclare(queueName, false, false, false, false, nil)

	fmt.Println("Consuming on: " + queueName)

	//channel.QueueBind(queueName, "#", "", false, nil)

	if err != nil {
		fmt.Println(err)
	}

	responses, err := channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}

	err = channel.Publish(
		"data-provision",
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: correlationId,
			ReplyTo:       q.Name,
			Body:          []byte(message),
		},
	)

	if err != nil {
		log.Fatalln("Could not publish " + message + " to queue: " + err.Error())
	}

	for res := range responses {
		if correlationId == res.CorrelationId {
			payload := string(res.Body)
			channel.QueueDelete(queueName, false, false, false)
			c <- payload
			break
		}
	}
}
