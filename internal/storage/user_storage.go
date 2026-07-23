package storage

import (
	"encoding/json"
	"os"
	"practic/internal/config"
	"practic/internal/models"
)

func LoadUsers() ([]models.User, error) {
	fileData, err := os.ReadFile(config.UsersFile)
	if err != nil {
		return nil, err
	}

	var users []models.User
	if err := json.Unmarshal(fileData, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func SaveUsers(users []models.User) error {
	usersJson, err := json.MarshalIndent(users, "", "	")
	if err != nil {
		return err
	}

	return os.WriteFile(config.UsersFile, usersJson, 0644)
}

func FindUserByEmail(users []models.User, email string) bool {
	for _, user := range users {
		if user.Email == email {
			return true
		}
	}

	return false
}

func FindUserByID(users []models.User, id string) bool {
	for _, user := range users {
		if user.ID == id {
			return true
		}
	}

	return false
}
