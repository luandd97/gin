package publisher

import (
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

func joinNetwork() {
	conn, err := amqp.Dial(os.Getenv("AMQP_URL"))
	handleError(err, "Can't connect to AMQP")
	defer conn.Close()

	C, _ := conn.Channel()
	handleError(err, "Can't create a amqpChannel")
	defer C.Close()
	queue, err := C.QueueDeclare("TestQueue", false, false, false, false, nil)
	handleError(err, "Can't Queue")
	fmt.Println(queue)

	// err = C.Publish(
	// 	"",
	// 	"TestQueue",
	// 	false,
	// 	false,
	// 	amqp.Publishing{
	// 		ContentType: "text/plain",
	// 		Body:        []byte("Hello World"),
	// 	},
	// )
	// handleError(err, "Queue Failed")
}

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
