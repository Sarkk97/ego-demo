package main

import (
	"log"
	"os"
)

var (
	//RabbitHost is RabbitMQ Host
	RabbitHost string
	//RabbitPort is RabbitMQ Port
	RabbitPort string
	//RabbitUsername is RabbitMQ Host
	RabbitUsername string
	//RabbitPassword is RabbitMQ Password
	RabbitPassword string

	exist bool
)

func init() {
	if RabbitHost, exist = os.LookupEnv("RABBIT_HOST"); !exist {
		log.Panicln("RABBIT_HOST env variable not set")
	}

	if RabbitPort, exist = os.LookupEnv("RABBIT_PORT"); !exist {
		log.Panicln("RABBIT_PORT env variable not set")
	}

	if RabbitUsername, exist = os.LookupEnv("RABBIT_USERNAME"); !exist {
		log.Panicln("RABBIT_USERNAME env variable not set")
	}

	if RabbitPassword, exist = os.LookupEnv("RABBIT_PASSWORD"); !exist {
		log.Panicln("RABBIT_PASSWORD env variable not set")
	}
}

func main() {
	runSubscriber("email_notifications", sendMail) //Spin up a queue worker for sending emails
}
