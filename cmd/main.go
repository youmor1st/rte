package main

import (
	"finalgo/routes"
	"finalgo/utils"
	"fmt"
	"net/http"
)

func main() {
	err := utils.InitDB()
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}

	routes.SetupRoutes()

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}
