package app

import (
	"fmt"
	"net/http"
	"os"

	"example.com/app/routes"
)

func NewServer() {
	fmt.Println("Server running on PORT " + os.Getenv("APP_PORT"))
	http.ListenAndServe(":"+os.Getenv("APP_PORT"), routes.Router())
}
