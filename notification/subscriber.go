package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

//runScriber spins up a queue worker for a queue
func runSubscriber(queueName string, fn func(message []byte)) {

	conn, err := amqp.Dial(fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		RabbitUsername,
		RabbitPassword,
		RabbitHost,
		RabbitPort,
	))
	failOnError(err, "Failed to connect ot Rabbit")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"notifications", //name
		"topic",         //type
		true,            //durable
		false,           //auto-deleted
		false,           //internal
		false,           //no-wait
		nil,             //arguments
	)
	failOnError(err, "Failed to define a queue")

	_, err = ch.QueueDeclare(
		queueName, //name
		true,      //durable
		false,     //delete when unused
		false,     //exclusive
		false,     //no-wait
		nil,       //arguments
	)
	failOnError(err, "Failed to declare queue user_mgt")

	err = ch.QueueBind(
		queueName,       //queue name
		queueName,       //routing key/binding key
		"notifications", // Exchange name
		false,
		nil,
	)
	failOnError(err, "Failed to bind to queue")

	msgs, err := ch.Consume(
		queueName, // queue
		"",        // consumer
		false,     // auto ack
		false,     // exclusive
		false,     // no local
		false,     // no wait
		nil,       // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			fn(d.Body)
			d.Ack(false)
		}
	}()

	<-forever
}

func sendMail(message []byte) {
	email := &struct {
		From    string   `json:"from"`
		To      []string `json:"to"`
		Bcc     []string `json:"bcc"`
		Cc      []string `json:"cc"`
		Subject string   `json:"subject"`
		Body    []byte   `json:"body"`
	}{}

	json.Unmarshal(message, email)

	smtpMailer := NewMailer(
		email.From,
		email.To,
		email.Bcc,
		email.Cc,
		email.Subject,
		email.Body,
	)

	smtpMailer.Send()
}
