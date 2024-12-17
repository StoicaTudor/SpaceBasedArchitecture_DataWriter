package repository

import (
	"DataWriter/data"
	"DataWriter/database_connector"
	"errors"
	"log"
)

func CreateUser(user data.User) error {
	logErrorMessage := "Error creating user: %v"
	appErrorMessage := "cannot create user"

	db, dbInitError := database_connector.InitDB()
	if dbInitError != nil {
		log.Printf(logErrorMessage, dbInitError)
		return errors.New(appErrorMessage)
	}
	defer database_connector.CloseDB()

	query := "INSERT INTO user (id, name, balance) VALUES (?, ?, ?)"
	_, queryExecError := db.Exec(query, user.ID, user.Name, user.Balance)
	if queryExecError != nil {
		log.Printf("Error creating user: %v", dbInitError)
		return dbInitError
	}
	return nil
}

func UpdateUser(user data.User) error {
	logErrorMessage := "Error updating user: %v"
	appErrorMessage := "cannot update user"

	db, dbInitError := database_connector.InitDB()
	if dbInitError != nil {
		log.Printf(logErrorMessage, dbInitError)
		return errors.New(appErrorMessage)
	}
	defer database_connector.CloseDB()

	query := "UPDATE user SET name = ?, balance = ? WHERE id = ?"
	_, err := db.Exec(query, user.Name, user.Balance, user.ID)
	if err != nil {
		log.Printf("Error updating user: %v", err)
		return err
	}
	return nil
}

func DeleteUser(id string) error {
	logErrorMessage := "Error deleting user: %v"
	appErrorMessage := "cannot deleting user"

	db, dbInitError := database_connector.InitDB()
	if dbInitError != nil {
		log.Printf(logErrorMessage, dbInitError)
		return errors.New(appErrorMessage)
	}
	defer database_connector.CloseDB()

	query := "DELETE FROM user WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting user: %v", err)
		return err
	}
	return nil
}

//// GetUser retrieves a user by their ID
//func GetUser(db *sql.DB, id string) (*data.User, error) {
//	query := "SELECT id, name, balance FROM user WHERE id = ?"
//	row := db.QueryRow(query, id)
//
//	var user data.User
//	err := row.Scan(&user.ID, &user.Name, &user.Balance)
//	if err != nil {
//		if err == sql.ErrNoRows {
//			return nil, fmt.Errorf("user with ID %s not found", id)
//		}
//		log.Printf("Error retrieving user: %v", err)
//		return nil, err
//	}
//	return &user, nil
//}
