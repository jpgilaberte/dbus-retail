package rabbitmq

import (
	"github.com/streadway/amqp"
	"github.com/Sirupsen/logrus"
	"github.com/jpgilaberte/dbus-retail/routes"
	"context"
)

const()
var()

func Run(ctx context.Context) {
	conn, err := amqp.Dial("amqp://admin_usuario:admlapasword@192.168.121.75:5672/")
	if err != nil {
		logrus.Fatalf("Failed to connect to RabbitMQ: %v", err)
		return
	}

	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		logrus.Fatal("Failed to open a channel")
		return
	}

	defer ch.Close()
	q, err := ch.QueueDeclare(
		"rpc_queue", // name
		false,       // durable
		false,       // delete when usused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		logrus.Fatal("Failed to declare a queue")
		return
	}

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		logrus.Fatal("Failed to set QoS")
		return
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		logrus.Fatal("Failed to register a consumer")
		return
	}

	for d := range msgs {
		go routes.Route(d.Body, rpcResponseFunc(ch, d))
	}
}

func rpcResponseFunc(ch *amqp.Channel, d amqp.Delivery) func(response []byte){
	dAux := d
	return func (response []byte){
		err := ch.Publish(
			"",        // exchange
			dAux.ReplyTo, // routing key
			false,     // mandatory
			false,     // immediate
			amqp.Publishing{
				ContentType:   "text/json",
				CorrelationId: dAux.CorrelationId,
				Body:          response,
			})
		if err != nil {
			logrus.Fatal("Failed to publish a message")

		}
		dAux.Ack(false)
	}
}

