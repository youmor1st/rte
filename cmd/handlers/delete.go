package handlers

import (
	utils2 "finalgo/pkg/models"
	"html/template"
	"net/http"
	"strconv"
)

func DeleteSelectHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/delete_select.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {
		userType := r.FormValue("user_type")

		if userType == "student" {
			http.Redirect(w, r, "/admin/delete-select-class", http.StatusSeeOther)
		} else if userType == "teacher" {
			http.Redirect(w, r, "/admin/delete-teacher", http.StatusSeeOther)
		} else {
			http.Error(w, "Invalid user type", http.StatusBadRequest)
		}

		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func DeleteSelectClassHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/select_class.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {
		classID := r.FormValue("class_id")
		http.Redirect(w, r, "/admin/delete-student?class_id="+classID, http.StatusSeeOther)
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

func DeleteStudentHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/delete_student.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Получаем идентификатор класса из формы
	classID := r.FormValue("class_id")
	if classID == "" {
		http.Error(w, "Missing class_id parameter", http.StatusBadRequest)
		return
	}

	classIDInt, err := strconv.Atoi(classID)
	if err != nil {
		http.Error(w, "Invalid class_id parameter: "+classID, http.StatusBadRequest)
		return
	}

	students, err := utils2.GetStudentsByClassID(classIDInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {
		studentID := r.FormValue("student_id")
		studentIDInt, err := strconv.Atoi(studentID)
		if err != nil {
			http.Error(w, "Invalid student_id parameter: "+studentID, http.StatusBadRequest)
			return
		}

		err = utils2.DeleteUser(studentIDInt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/admin/delete-student?class_id="+classID, http.StatusSeeOther)
		return
	}

	data := struct {
		ClassID  int
		Students []utils2.User
	}{
		ClassID:  classIDInt,
		Students: students,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
