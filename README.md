# PubSub Example

## 1. Install all required dependencies
  Go and its dependencies
  
  RabbitMQ
  
  MySql

## 2. Create Database
 Create database 'pubsub'
## 3. Start Receiver

    go run PubSubExample/receiver/main.go 

## 4. Publish data

    go run PubSubExample/sender/main.go 
    
## 5. Test
  Open database and check data
