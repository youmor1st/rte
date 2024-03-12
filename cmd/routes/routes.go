// routes/routes.go

package routes

import (
	handlers2 "finalgo/cmd/handlers"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/login", handlers2.LoginHandler)

	http.HandleFunc("/admin/student-register", handlers2.RegistrationHandler)

	http.HandleFunc("/admin/delete-select", handlers2.DeleteSelectHandler)
	http.HandleFunc("/admin/delete-select-class", handlers2.DeleteSelectClassHandler)
	http.HandleFunc("/admin/delete-student", handlers2.DeleteStudentHandler)

	http.HandleFunc("/admin/update-select", handlers2.UpdateSelectHandler)
	http.HandleFunc("/admin/update-select-student", handlers2.UpdateSelectStudentHandler)
	http.HandleFunc("/admin/update-student", handlers2.UpdateStudentHandler)

	http.HandleFunc("/admin/users", handlers2.UsersHandler)

}
