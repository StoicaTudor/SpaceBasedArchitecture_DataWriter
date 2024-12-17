package adapter

import (
	"DataWriter/data_supply/dtos"
	"encoding/json"
)

type DeserializedCommand struct {
	User json.RawMessage `json:"user"`
	Type dtos.ActionType `json:"type"`
}

func ConvertToCommand(serializedCommand []byte) dtos.Command {
	var deserializedCommand DeserializedCommand
	error0 := json.Unmarshal(serializedCommand, &deserializedCommand)

	if error0 != nil {
		panic("bazdmeg0")
	}

	switch deserializedCommand.Type {
	case dtos.CREATE:
		var dto dtos.UserCreateDTO
		error3 := json.Unmarshal(deserializedCommand.User, &dto)
		//command, error1 := util.Cast[data_contracts.User, data_contracts.UserCreateDTO](deserializedCommand.User)
		if error3 != nil {
			panic("bazdmeg1")
		}
		return dto

	case dtos.UPDATE:
		var dto dtos.UserUpdateDTO
		error3 := json.Unmarshal(deserializedCommand.User, &dto)
		if error3 != nil {
			panic("bazdmeg1")
		}
		return dto

	case dtos.DELETE:
		var dto dtos.UserDeleteDTO
		error3 := json.Unmarshal(deserializedCommand.User, &dto)
		if error3 != nil {
			panic("bazdmeg1")
		}
		return dto

	default:
		panic("bazdmeg2")
	}
}
