// routes/routes.go

package routes

import (
	"finalgo/handlers"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/admin/student-register", handlers.RegistrationHandler)
	http.HandleFunc("/admin/delete-select", handlers.DeleteSelectHandler)
	http.HandleFunc("/admin/delete-select-class", handlers.DeleteSelectClassHandler)
	http.HandleFunc("/admin/delete-student", handlers.DeleteStudentHandler)

	// Добавьте другие маршруты при необходимости
}
