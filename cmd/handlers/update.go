package handlers

import (
	utils2 "finalgo/pkg/models"
	"html/template"
	"net/http"
	"strconv"
)

func UpdateSelectHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/update_select.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {
		classID := r.FormValue("class_id")
		http.Redirect(w, r, "/admin/update-select-student?class_id="+classID, http.StatusSeeOther)
		return
	}

	classes, err := utils2.GetClasses()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Classes []utils2.Class
	}{
		Classes: classes,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func UpdateSelectStudentHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/update_select_student.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	classID := r.URL.Query().Get("class_id")

	if classID == "" {
		http.Error(w, "Missing class_id parameter", http.StatusBadRequest)
		return
	}

	classIDInt, err := strconv.Atoi(classID)
	if err != nil {
		http.Error(w, "Invalid class_id parameter", http.StatusBadRequest)
		return
	}

	// Получаем учеников из базы данных для данного класса
	students, err := utils2.GetStudentsByClassID(classIDInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		ClassID  string
		Students []utils2.User
	}{
		ClassID:  classID,
		Students: students,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UpdateStudentHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/update_student.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {
		studentID := r.FormValue("student_id")
		username := r.FormValue("username")
		password := r.FormValue("password")
		firstName := r.FormValue("firstName")
		lastName := r.FormValue("lastName")
		classID := r.FormValue("classID")
		points := r.FormValue("points")

		studentIDInt, _ := strconv.Atoi(studentID)
		classIDInt, _ := strconv.Atoi(classID)
		pointsInt, _ := strconv.Atoi(points)

		updatedStudent := utils2.User{
			ID:       studentIDInt,
			Username: username,
			Password: password,
			FName:    firstName,
			SName:    lastName,
			ClassID:  classIDInt,
			Points:   pointsInt,
		}

		err := utils2.UpdateUser(updatedStudent)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/admin/update-select", http.StatusSeeOther)
		return
	}

	studentID := r.URL.Query().Get("student_id")
	if studentID == "" {
		http.Error(w, "Missing student_id parameter", http.StatusBadRequest)
		return
	}

	studentIDInt, err := strconv.Atoi(studentID)
	if err != nil {
		http.Error(w, "Invalid student_id parameter", http.StatusBadRequest)
		return
	}

	student, err := utils2.GetStudentByID(studentIDInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Student utils2.User
	}{
		Student: student,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
