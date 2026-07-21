package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"practic/internal/models"
	"practic/internal/storage"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid json", http.StatusBadRequest)
		fmt.Println("Invalid json")
		return
	}

	users, err := storage.LoadUsers()
	if err != nil {
		http.Error(w, "Error load users", http.StatusInsufficientStorage)
		return
	}

	if storage.FindUserByEmail(users, user.Email) {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	newUser := models.NewUser(user.Email)

	users = append(users, newUser)
	if err := storage.SaveUsers(users); err != nil {
		http.Error(w, "Erorr save users", http.StatusInternalServerError)
		return
	}

	if err := storage.CreateUserFolder(newUser.ID); err != nil {
		http.Error(w, "Error create user file", http.StatusInternalServerError)
	}

	var response models.RegisterResponse
	response.ID = newUser.ID

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Println(err)
	}
}
