package dtos

import "DataWriter/data"

type UserUpdateDTO struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

func (dto UserUpdateDTO) GetAction() ActionType {
	return UPDATE
}

func (dto UserUpdateDTO) GetCommandType() CommandType {
	return UpdateUserDto
}

func (dto UserUpdateDTO) ToUser() data.User {
	return data.User{
		ID:      dto.ID,
		Name:    dto.Name,
		Balance: dto.Balance,
	}
}
