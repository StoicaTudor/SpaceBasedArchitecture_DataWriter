package environment

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const BaseEnvironmentFilePath string = "./environment/.env"
const DevelopmentEnvironmentFilePath string = "./environment/.development-environment"
const ProductionEnvironmentFilePath string = "./environment/.production-environment"

type name string

const (
	Development name = "development"
	Production  name = "production"
)

type VariableName string

const (
	Environment  VariableName = "ENVIRONMENT"
	KafkaBroker  VariableName = "KAFKA_BROKER"
	KafkaTopic   VariableName = "KAFKA_TOPIC"
	KafkaGroupID VariableName = "KAFKA_GROUP_ID"
	DBUser       VariableName = "DB_USER"
	DBPassword   VariableName = "DB_PASSWORD"
	DBHost       VariableName = "DB_HOST"
	DBPort       VariableName = "DB_PORT"
	DBName       VariableName = "DB_NAME"
	MongoURI     VariableName = "MONGO_URI"
	MongoDB      VariableName = "MONGO_DB"
	MongoColl    VariableName = "MONGO_COLL"

	DataSupplySource VariableName = "DATA_SUPPLY_SOURCE"
)

/*
LEARN: return type -> error
The error built-in interface type is the conventional interface for representing an error condition,
with the nil value representing no error.
*/
func loadEnv() error {
	baseEnvironment, baseEnvironmentLoadingError := loadBaseEnvironment()
	if baseEnvironmentLoadingError != nil {
		return baseEnvironmentLoadingError
	}

	correspondingEnvironment, correspondingEnvironmentLoadingError := loadCorrespondingEnvironment(baseEnvironment)
	if correspondingEnvironmentLoadingError != nil {
		return correspondingEnvironmentLoadingError
	}

	// Load the environment-specific file
	if err := godotenv.Load(correspondingEnvironment); err != nil {
		return fmt.Errorf("error loading environment file %s: %v", correspondingEnvironment, err)
	}

	log.Printf("Loaded environment configuration from %s", correspondingEnvironment)
	return nil
}

func loadCorrespondingEnvironment(baseEnvironment string) (string, error) {
	// Load the corresponding environment file
	var envFile string
	switch baseEnvironment {
	case string(Development):
		envFile = DevelopmentEnvironmentFilePath
	case string(Production):
		envFile = ProductionEnvironmentFilePath
	default:
		return "", fmt.Errorf("invalid ENVIRONMENT value: %s", baseEnvironment)
	}
	return envFile, nil
}

// LEARN: you can return multiple things in the same shot in GO
// LEARN: you can return perform if with preparation in GO
func loadBaseEnvironment() (string, error) {
	// Load the main .env file to get the ENVIRONMENT variable
	if err := godotenv.Load(BaseEnvironmentFilePath); err != nil {
		return "", fmt.Errorf("error loading .env file: %v", err)
	}

	// Get the ENVIRONMENT variable
	environment := os.Getenv(string(Environment))
	if environment == "" {
		return "", fmt.Errorf("ENVIRONMENT variable is not set in .env")
	}
	return environment, nil
}

func Load() {
	// Load the environment files
	if err := loadEnv(); err != nil {
		log.Fatalf("Failed to load environment files: %v", err)
	}

	// Continue with the rest of your application
	log.Println("Environment loaded successfully!")

	// For example, print out a variable to verify that it was loaded correctly
	kafkaBroker := os.Getenv(string(KafkaBroker))
	log.Printf("Kafka broker: %s", kafkaBroker)
}
