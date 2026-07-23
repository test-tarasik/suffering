package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"practic/internal/models"
	"practic/internal/storage"
)

func UploadFiles(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	id := r.FormValue("user_id")

	users, err := storage.LoadUsers()
	if err != nil {
		http.Error(w, "Error load users", http.StatusInternalServerError)
		return
	}

	if !storage.FindUserByID(users, id) {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Invalid file", http.StatusBadRequest)
		return
	}

	defer file.Close()

	if err := storage.SaveFile(id, header.Filename, file); err != nil {
		http.Error(w, "Error save file", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)

}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid json", http.StatusBadRequest)
		fmt.Println("Invalid json")
		return
	}

	users, err := storage.LoadUsers()
	if err != nil {
		http.Error(w, "Error load users", http.StatusInternalServerError)
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
		return
	}

	var response models.RegisterResponse
	response.ID = newUser.ID

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Println(err)
	}
}
