package controller

import (
	"database/sql"
	"fmt"
	"go-sample/model"
	"net/http"
)

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	users, err := model.GetAllUsers(db)
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	for _, user := range users {
		fmt.Fprintf(w, "ID: %d, Username: %s, Email: %s\n", user.ID, user.Username, user.Email)
	}
}


test