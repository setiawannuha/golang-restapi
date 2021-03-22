package routes

import (
	"net/http"

	"example.com/app/configs/db"
	"example.com/app/controllers"
	"example.com/app/helpers"
	"example.com/app/middlewares"
	"github.com/gorilla/mux"
)

func Router() http.Handler {
	db, err := db.Postgres()
	if err != nil {
		helpers.Log("Failed to load database")
	}
	user := controllers.Config{DB: db, Driver: "postgresql"}

	r := mux.NewRouter().StrictSlash(false)
	api := r.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/ping", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("PONG!"))
	})

	r.HandleFunc("/register", user.InsertUser).Methods("POST")
	r.HandleFunc("/signin", user.Signin).Methods("POST")

	api.HandleFunc("/users", user.GetAllUsers).Methods("GET")
	api.HandleFunc("/users/{id}", user.GetDetailUser).Methods("GET")

	// Middleware
	api.Use(middlewares.LogMiddleware)
	api.Use(middlewares.Authorization)
	// api.Use(func(h http.Handler) http.Handler {
	// 	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	// 		fmt.Println("Middleware")
	// 		h.ServeHTTP(rw, r)
	// 	})
	// })
	return r
}
