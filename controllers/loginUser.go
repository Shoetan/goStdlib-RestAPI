package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/shoetan/ecom/types"
	"github.com/shoetan/ecom/utils"
)

func LoginUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get body of request
		var payload types.LoginPayload
		json.NewDecoder(r.Body).Decode(&payload)

		// Query the database for the user with the provided email
		row := db.QueryRow("SELECT email, password FROM users WHERE email = $1", payload.Email)

		var email, password string
		err := row.Scan(&email, &password)
		if err != nil {
			if err == sql.ErrNoRows {
				// No user found with the provided email
				http.Error(w, "User not found", http.StatusNotFound)
				return
			}
			// Error while querying the database
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		comparePwd := utils.ComparePassword(password, payload.Password)

		if comparePwd != nil {
			http.Error(w, "Wrong Password",http.StatusBadRequest)
		} else {
			token, tokenErr := utils.CreateToken(payload.Email)

			userCred := types.LoginResponse{
				Email:    email,
				Token: token,
			}

			if tokenErr != nil {
				http.Error(w,"Invalid Token",http.StatusInternalServerError)
			} else {
				err = json.NewEncoder(w).Encode(userCred)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			}
		}
	}
}
