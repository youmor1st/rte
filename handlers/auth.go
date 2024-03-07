// В вашем файле handlers/handlers.go

package handlers

import (
	"finalgo/utils"
	"fmt"
	"html/template"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {
		// Обработка входа
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Проверка логина и пароля
		validUser, _ := utils.ValidateUser(username, password)

		if validUser {
			fmt.Printf("Login successful for user: %s\n", username)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		} else {
			fmt.Println("Invalid login attempt")
		}
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
