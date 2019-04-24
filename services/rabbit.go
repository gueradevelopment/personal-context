package services

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"log"
)

type RabbitService struct {
	connection *amqp.Connection
	channel    *amqp.Channel
}

func RabbitServiceInit() *RabbitService {
	service := new(RabbitService)
	connection, err := amqp.Dial("amqp://guera:pass@34.73.163.109")
	if err != nil {
		log.Fatalln("Rabbit Connection Error")
		return nil
	}
	service.connection = connection
	channel, err := service.connection.Channel()
	if err != nil {
		log.Fatalln("Channel Creation Error")
		return nil
	}
	service.channel = channel
	return service
}

func (service *RabbitService) SendAndReceive(message string, routingKey string, c chan string) {
	defer close(c)
	service.channel.ExchangeDeclare(
		"data-provision",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	correlationId := uuid.New().String()
	queueName := uuid.New().String()
	q, err := service.channel.QueueDeclare(queueName, false, true, false, true, nil)
	service.channel.QueueBind(queueName, "data-provision", "#", false, nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	//if service.channel.closed
	err = service.channel.Publish(
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
	responses, _ := service.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	for res := range responses {
		if correlationId != res.CorrelationId {
			continue
		}
		payload := string(res.Body)
		fmt.Println(payload)
		service.channel.QueueDelete(q.Name, false, false, false)
		c <- payload
	}
}
