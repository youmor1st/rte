package handlers

import (
	"finalgo/pkg/models"
	"html/template"
	"net/http"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	usersWithClass, err := models.GetAllUsersWithClass()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl, err := template.ParseFiles("web/templates/users.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, usersWithClass)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
