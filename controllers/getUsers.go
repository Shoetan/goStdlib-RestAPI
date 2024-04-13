package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/shoetan/ecom/types"
)

type UserDetails struct {
	Id uint
	Name string
	Email string
}



func GetAllUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM users")

		if err != nil {
			fmt.Println("Error:", err)
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
		}
		defer rows.Close()

		users := []UserDetails{}

		for rows.Next(){
			var u types.User

			if err := rows.Scan(&u.Id, &u.Name, &u.Email, &u.Password); err!=nil {
				log.Fatal(err)
			}

			userResponse := UserDetails{
				Id: u.Id,
				Name: u.Name,
				Email: u.Email,
				
			}

			users = append(users, userResponse)

		}

		err = json.NewEncoder(w).Encode(users)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}