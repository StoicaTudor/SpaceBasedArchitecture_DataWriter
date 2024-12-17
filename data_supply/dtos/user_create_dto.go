package dtos

import (
	"DataWriter/data"
	"fmt"
	"github.com/google/uuid"
	"log"
	"math/rand"
	"time"
)

type UserCreateDTO struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

func (dto UserCreateDTO) GetAction() ActionType {
	return CREATE
}

func (dto UserCreateDTO) GetCommandType() CommandType {
	return CreateUserDto
}

func (dto UserCreateDTO) ToUser() data.User {
	return data.User{
		ID:      dto.ID,
		Name:    dto.Name,
		Balance: dto.Balance,
	}
}

func GetRandomCommand() Command {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	switch random.Intn(3) {
	case 0:
		return &UserCreateDTO{
			ID:      uuid.New().String(),
			Name:    fmt.Sprintf("User%d", random.Intn(100)),
			Balance: rand.Float64() * 1000,
		}
	case 1:
		return &UserUpdateDTO{
			ID:      uuid.New().String(),
			Name:    fmt.Sprintf("User%d", random.Intn(100)),
			Balance: random.Float64() * 1000,
		}
	case 2:
		return &UserDeleteDTO{
			ID: uuid.New().String(),
		}
	default:
		return nil
	}
}

// Function to handle the command and cast it to UserCreateDTO
func handleCommand(command Command) {
	// Check if the command type is CREATE_USER_DTO
	if command.GetCommandType() == CreateUserDto {
		// Perform type assertion (cast)
		if userDTO, ok := command.(*UserCreateDTO); ok {
			// Successfully casted to UserCreateDTO
			fmt.Printf("UserCreateDTO: %+v\n", userDTO)
		} else {
			log.Fatalf("Failed to cast command to UserCreateDTO")
		}
	} else {
		log.Printf("Unhandled command type: %s", command.GetCommandType())
	}
}
