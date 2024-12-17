package service

import (
	"DataWriter/data_supply/dtos"
	"DataWriter/repository"
)

// TODO: handle errors
func CreateUser(dto dtos.UserCreateDTO) {
	repository.CreateUser(dto.ToUser())
}

func UpdateUser(dto dtos.UserUpdateDTO) {
	repository.UpdateUser(dto.ToUser())
}

func DeleteUser(dto dtos.UserDeleteDTO) {
	repository.DeleteUser(dto.ID)
}
