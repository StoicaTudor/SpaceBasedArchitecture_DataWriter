package dtos

type UserDeleteDTO struct {
	ID string `json:"id"`
}

func (dto UserDeleteDTO) GetAction() ActionType {
	return DELETE
}

func (dto UserDeleteDTO) GetCommandType() CommandType {
	return DeleteUserDto
}
