package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"fmt"

	"github.com/shoetan/ecom/types"
	"github.com/shoetan/ecom/utils"
)

type UserResponse struct {
	ID int `json:"id"`
	Email string `json:"email"`
	Name string `json:"name"`
}

func CreateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user types.User


		json.NewDecoder(r.Body).Decode(&user)

		hashPwd, err := utils.HashPwd(user.Password)

		if err != nil {
			fmt.Println(err.Error())
		}

		user.Password = hashPwd

		_, err = db.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", user.Name, user.Email, user.Password)

		if err != nil {
			fmt.Println("Error:", err)
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return
		}

		userResponse := UserResponse {
			ID: int(user.Id),
			Email: user.Email,
			Name: user.Name,

		}

		json.NewEncoder(w).Encode(userResponse)
	}
}
