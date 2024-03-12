package handlers

import (
	"finalgo/utils"
	"html/template"
	"net/http"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	usersWithClass, err := utils.GetAllUsersWithClass()
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
