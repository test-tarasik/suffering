package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/google/uuid"
)

type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type RegisterResponse struct {
	ID string `json:"id"`
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser User

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Invalid json", http.StatusBadRequest)
		fmt.Println("Invalid json")
		return
	}

	fileData, err := os.ReadFile("data/users.json")
	if err != nil {
		fmt.Println("err:", err)
	}

	var users []User
	json.Unmarshal(fileData, &users)

	for _, v := range users {
		if v.Email == newUser.Email {
			http.Error(w, "Такой email уже существует", http.StatusConflict)
			return
		}
	}

	id := uuid.New().String()

	newUser.ID = id

	users = append(users, newUser)

	usersJson, err := json.MarshalIndent(users, "", "	")
	if err != nil {
		fmt.Println("Error encode slice to json")
		http.Error(w, "Ошибка преобразования слайса", http.StatusInternalServerError)
		return
	}

	if err := os.WriteFile("data/users.json", usersJson, 0644); err != nil {
		http.Error(w, "error1", http.StatusInternalServerError)
		return
	}

	if err := os.Mkdir("files/"+newUser.ID, 0755); err != nil {
		http.Error(w, "error2", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	var response RegisterResponse
	response.ID = newUser.ID

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Println("err encode:", err)
	}
}
