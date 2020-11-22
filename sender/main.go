package main

import (
	"../utils/configHelper"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func init() {
	configHelper.InitViper()
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Errorf("%s: %s", msg, err)
	}
}

func main() {
	userName := configHelper.GetConfig("rabbitMq.userName")
	password := configHelper.GetConfig("rabbitMq.password")
	host := configHelper.GetConfig("rabbitMq.host")
	port := configHelper.GetConfig("rabbitMq.port")

	conn, err := amqp.Dial("amqp://" + userName + ":" + password + "@" + host + ":" + port + "/")
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

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(msg),
		})
	failOnError(err, "Failed to publish a message")
}

var msg = `{
	"offers": [
	{
	"cm_offer_id": "8f6995366e854c9faf1d9f3d233702b8",
	"hotel": {
	"hotel_id": "BH~46456",
	"name": "Hawthorn Suites by Wyndham Eagle CO",
	"country": "US",
	"address": "0315 Chambers Avenue, 81631",
	"latitude": 39.660193,
	"longitude": -106.824123,
	"telephone": "+1-970-3283000",
	"amenities": [
	"Business Centre",
	"Fitness Room/Gym",
	"Pet Friendly",
	"Disabled Access",
	"Air Conditioned",
	"Free WIFI",
	"Elevator / Lift",
	"Parking"
	],
	"description": "Stay a while in beautiful mountain country at this Hawthorn Suites by Wyndham Eagle CO hotel, just off Interstate 70, only 6 miles from the Vail/Eagle Airport and close to skiing, golfing, Eagle River and great restaurants. Pets are welcome at this h",
	"room_count": 1,
	"currency": "USD"
	},
	"room": {
   "hotel_id": "BH~46456",
	"room_id": "S2Q",
	"description": "JUNIOR SUITES WITH 2 QUEEN BEDS",
	"name": "JUNIOR SUITES WITH 2 QUEEN BEDS",
	"capacity": {
	"max_adults": 2,
	"extra_children": 2
	}
	},
	"rate_plan": {
	"hotel_id": "BH~46456",
	"rate_plan_id": "BAR",
	"cancellation_policy": [
	{
	"type": "Free cancellation",
	"expires_days_before": 2
	}
	],
	"name": "BEST AVAILABLE RATE",
	"other_conditions": [
	"CXL BY 2 DAYS PRIOR TO ARRIVAL-FEE 1 NIGHT 2 DAYS PRIOR TO ARRIVAL",
	"BEST AVAILABLE RATE"
	],
	"meal_plan": "Room only"
	},
	"original_data": {
	"GuaranteePolicy": {
	"Required": true
	}
	},
	"capacity": {
	"max_adults": 2,
	"extra_children": 2
	},
	"number": 1,
	"price": 1520,
	"currency": "USD",
	"check_in": "2020-11-18",
	"check_out": "2020-11-20",
	"fees": [
	{
	"type": "CountyTax",
	"description": "COUNTY TAX PER STAY",
	"included": true,
	"percent": 17.5
	}
	]
	}
	]
   }`
