package services

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

type QueueService interface {
	MakePublish(routingKey string, data []byte, headers map[string]interface{})
	Init() amqp.Connection
}

type queueService struct{}

func NewQueueService() QueueService {
	return &queueService{}
}

func (s *queueService) Init() amqp.Connection {
	conn, err := amqp.Dial(os.Getenv("AMQP_URL"))
	handleError(err, "Can't connect to AMQP")
	return *conn
}

func (s *queueService) MakePublish(routingKey string, data []byte, headers map[string]interface{}) {
	godotenv.Load()
	amqpServer := os.Getenv("AMQP_URL")
	conn, err := amqp.Dial(amqpServer)
	handleError(err, "Can't connect to AMQP : "+amqpServer)
	defer conn.Close()

	C, err := conn.Channel()
	handleError(err, "Can't create a amqpChannel")
	defer C.Close()
	queue, err := C.QueueDeclare(routingKey, false, false, false, false, nil)
	handleError(err, "Can't Queue")
	fmt.Println(queue)
	err = C.Publish(
		"",
		routingKey,
		false,
		false,
		amqp.Publishing{
			Headers:         headers,
			ContentType:     "application/json",
			ContentEncoding: "",
			DeliveryMode:    0,
			Priority:        0,
			CorrelationId:   "",
			ReplyTo:         "",
			Expiration:      "",
			MessageId:       "",
			Timestamp:       time.Time{},
			Type:            "",
			UserId:          "",
			AppId:           "",
			Body:            data,
		},
	)
	fmt.Println(err)
	if err != nil {
		panic(err)
	}
}

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
