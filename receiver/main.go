package main

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"

	"../receiver/dao"
	"../receiver/model"
	"../receiver/service"
	"../utils/configHelper"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Errorf("%s: %s", msg, err)
	}
}

func init() {
	configHelper.InitViper()
	mySqlUserName := configHelper.GetConfig("mySql.userName")
	mySqlPassword := configHelper.GetConfig("mySql.password")
	mySqlHost := configHelper.GetConfig("mySql.host")
	mySqlPort := configHelper.GetConfig("mySql.port")
	mySqlDatabaseName := configHelper.GetConfig("mySql.databaseName")
	dao.InitDB(mySqlUserName, mySqlPassword, mySqlHost, mySqlPort, mySqlDatabaseName)
}

func main() {
	rabbitMqUserName := configHelper.GetConfig("rabbitMq.userName")
	rabbitMqPassword := configHelper.GetConfig("rabbitMq.password")
	rabbitMqHost := configHelper.GetConfig("rabbitMq.host")
	rabbitMqPort := configHelper.GetConfig("rabbitMq.port")

	conn, err := amqp.Dial("amqp://" + rabbitMqUserName + ":" + rabbitMqPassword + "@" + rabbitMqHost + ":" + rabbitMqPort + "/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"PubSubAssignment", // name
		false,              // durable
		false,              // delete when unused
		false,              // exclusive
		false,              // no-wait
		nil,                // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			msg := d.Body
			var myData model.DataHolder
			err := json.Unmarshal(msg, &myData)
			if err != nil {
				failOnError(err, "Fail to decode json data")
			}
			service.SaveOfferDataService(myData)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
