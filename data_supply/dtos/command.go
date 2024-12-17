package dtos

type ActionType string

const (
	CREATE ActionType = "CREATE"
	UPDATE ActionType = "UPDATE"
	DELETE ActionType = "DELETE"
)

type CommandType string

const (
	CreateUserDto CommandType = "CreateUserDTO"
	UpdateUserDto CommandType = "UpdateUserDTO"
	DeleteUserDto CommandType = "DeleteUserDTO"
)

type Command interface {
	GetAction() ActionType
	GetCommandType() CommandType
}
