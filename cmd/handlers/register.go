package handlers

import (
	utils2 "finalgo/pkg/models"
	"fmt"
	"net/http"
	"text/template"
)

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	isAdmin := true

	if !isAdmin {
		http.Error(w, "Access denied", http.StatusForbidden)
		return
	}

	tmpl, err := template.ParseFiles("web/templates/register.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	successfulRegistration := false

	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		fName := r.FormValue("f_name")
		sName := r.FormValue("s_name")
		className := r.FormValue("class_name")

		classID, err := utils2.GetOrCreateClass(className)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		user := utils2.User{
			Role:     "Student",
			Username: username,
			Password: password,
			FName:    fName,
			SName:    sName,
			ClassID:  classID,
			Points:   100,
		}

		err = utils2.CreateUser(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		successfulRegistration = true

		fmt.Println(w, "User registered successfully!")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	data := struct {
		SuccessfulRegistration bool
	}{
		SuccessfulRegistration: successfulRegistration,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
