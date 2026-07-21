package storage

import (
	"encoding/json"
	"os"
	"practic/internal/models"
)

const usersFile = "data/users.json"

func LoadUsers() ([]models.User, error) {
	fileData, err := os.ReadFile(usersFile)
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

	return os.WriteFile(usersFile, usersJson, 0644)
}

func FindUserByEmail(users []models.User, email string) bool {
	for _, user := range users {
		if user.Email == email {
			return true
		}
	}

	return false
}
