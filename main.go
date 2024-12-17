package main

import (
	"DataWriter/data_supplier_receiver"
	"DataWriter/data_supply/dtos"
	"DataWriter/environment"
	"DataWriter/service"
	"DataWriter/util"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func processMessage(message kafka.Message, mariaDB *sql.DB) {
	// Insert into MariaDB
	_, err := mariaDB.Exec("INSERT INTO users (data) VALUES (?)", string(message.Value))
	if err != nil {
		log.Printf("Error inserting into MariaDB: %v", err)
	}
}

func consumeKafka() {
	// Connect to MariaDB
	//mariaDB, err := connectMariaDB()
	//if err != nil {
	//	log.Fatalf("Error connecting to MariaDB: %v", err)
	//}
	//defer mariaDB.Close()

	// Kafka reader configuration
	kafkaReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{os.Getenv(string(environment.KafkaBroker))},
		Topic:   os.Getenv(string(environment.KafkaTopic)),
		GroupID: os.Getenv(string(environment.KafkaGroupID)),
	})

	// Consume messages from Kafka
	for {
		msg, err := kafkaReader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading message: %v", err)
			continue
		}

		// Process each message asynchronously
		//go processMessage(msg, mariaDB)
		go processMessage(msg, nil)
	}
}

func consumeKafka2() {
	// Get environment variables for Kafka configuration
	kafkaBroker := os.Getenv(string(environment.KafkaBroker))
	kafkaTopic := os.Getenv(string(environment.KafkaTopic))
	kafkaGroupID := os.Getenv(string(environment.KafkaGroupID))

	// Setup Kafka reader configuration
	readerConfig := kafka.ReaderConfig{
		Brokers: []string{kafkaBroker}, // Kafka broker addresses
		Topic:   kafkaTopic,            // Kafka topic
		GroupID: kafkaGroupID,          // Consumer group ID
	}

	// Create a new Kafka reader
	reader := kafka.NewReader(readerConfig)
	defer reader.Close()

	// Create a signal channel to gracefully handle termination (Ctrl+C)
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// Consume messages in a loop
	fmt.Println("Starting Kafka consumer...")

	for {
		// Read a message from Kafka
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading message: %v", err)
			continue
		}

		// Process the message (here we simply print it)
		fmt.Printf("Received message: %s\n", string(msg.Value))

		// Gracefully shut down on receiving SIGINT or SIGTERM
		select {
		case <-sigchan:
			fmt.Println("\nReceived termination signal, shutting down...")
			return
		default:
		}
	}
}

// Produce random strings to Kafka every 2 seconds
func produceRandomMessages() {
	// Kafka writer configuration
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{os.Getenv(string(environment.KafkaBroker))},
		Topic:   os.Getenv(string(environment.KafkaTopic)),
	})

	// Send a random message to Kafka every 2 seconds
	for {
		// Generate a random string
		randomString := util.GenerateRandomString(10)

		// Create Kafka message
		message := kafka.Message{
			Value: []byte(randomString),
		}

		// Write message to Kafka
		err := kafkaWriter.WriteMessages(context.Background(), message)
		if err != nil {
			log.Printf("Error writing message to Kafka: %v", err)
		} else {
			log.Printf("Produced message: %s", randomString)
		}

		// Wait for 2 seconds before producing the next message
		time.Sleep(2 * time.Second)
	}
}

type PrintConsumer struct{}

func (printConsumer *PrintConsumer) Consume(command dtos.Command) {
	switch castedCommand := command.(type) {
	case *dtos.UserCreateDTO:
		fmt.Println("UserCreateDTO - ID:", castedCommand.ID)
		fmt.Println("UserCreateDTO - Name:", castedCommand.Name)
		fmt.Println("UserCreateDTO - Balance:", castedCommand.Balance)
	case *dtos.UserUpdateDTO:
		fmt.Println("UserUpdateDTO - ID:", castedCommand.ID)
		fmt.Println("UserUpdateDTO - Name:", castedCommand.Name)
		fmt.Println("UserUpdateDTO - Balance:", castedCommand.Balance)
	case *dtos.UserDeleteDTO:
		fmt.Println("UserDeleteDTO - ID:", castedCommand.ID)
	default:
		fmt.Println("Unknown Command Type")
	}
}

type DBInteractionConsumer struct{}

func (consumer *DBInteractionConsumer) Consume(command dtos.Command) {
	switch command.GetCommandType() {
	case dtos.CreateUserDto:
		dto, err := util.Cast[dtos.Command, dtos.UserCreateDTO](command)
		if err != nil {
			return
		}
		service.CreateUser(dto)
		break

	case dtos.UpdateUserDto:
		dto, err := util.Cast[dtos.Command, dtos.UserUpdateDTO](command)
		if err != nil {
			return
		}
		service.UpdateUser(dto)
		break

	case dtos.DeleteUserDto:
		dto, err := util.Cast[dtos.Command, dtos.UserDeleteDTO](command)
		if err != nil {
			return
		}
		service.DeleteUser(dto)
		break
	}
}

func main() {
	// LEARN: export function from file -> start with uppercase
	environment.Load()

	// ---------------------

	// Start the Kafka producer in the background
	//go produceRandomMessages()

	// Start consuming messages from Kafka
	//consumeKafka2()
	// ---------------------

	// ------------------------------------------------------------------------------------------------------------------------------------------------
	//// Initialize the DB
	//db, err := repository.InitDB()
	//if err != nil {
	//	log.Fatalf("Error initializing database: %v", err)
	//	return
	//}
	//defer repository.CloseDB()
	//
	//// Create a new user
	//user := data_contracts.User{
	//	ID:      "1",
	//	Name:    "John Doe",
	//	Balance: 1000.50,
	//}
	//err = repository.CreateUser(db, user)
	//if err != nil {
	//	log.Fatalf("Error creating user: %v", err)
	//	return
	//}
	//fmt.Println("User created!")
	//
	//// Retrieve the user
	//retrievedUser, err := repository.GetUser(db, "1")
	//if err != nil {
	//	log.Fatalf("Error retrieving user: %v", err)
	//	return
	//}
	//fmt.Printf("Retrieved User: %+v\n", retrievedUser)
	//
	//// Update the user
	//retrievedUser.Balance += 500.00
	//err = repository.UpdateUser(db, *retrievedUser)
	//if err != nil {
	//	log.Fatalf("Error updating user: %v", err)
	//	return
	//}
	//fmt.Println("User updated!")
	// ------------------------------------------------------------------------------------------------------------------------------------------------
	//// Delete the user
	//err = repository.DeleteUser(db, "1")
	//if err != nil {
	//	log.Fatalf("Error deleting user: %v", err)
	//	return
	//}
	//fmt.Println("User deleted!")

	// ------------------
	//var wg sync.WaitGroup
	//
	//// LEARN: export function from file -> start with uppercase
	//environment.Load()
	//
	//// ---------------------
	////log.Println("Starting Kafka consumer...")
	////
	////// Start the Kafka producer in the background
	////go produceRandomMessages()
	////
	////// Start consuming messages from Kafka
	////consumeKafka()
	//// ---------------------
	//
	//dataSupplier, _ := data_supplier_receiver.GetDataSupplier()
	//wg.Add(1)
	//dataSupplier.Supply(&PrintConsumer{}, &wg)
	//wg.Wait()
	// ------------------

	dataSupplier, _ := data_supplier_receiver.GetDataSupplier()
	dataSupplier.Supply(&DBInteractionConsumer{}, nil)
}
