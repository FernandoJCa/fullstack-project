package main

import (
	_ "backend/docs"
	"backend/internal/database"
	"backend/internal/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

// @title Todo API
// @version 1.0
// @description This is a simple todo list API.
// @host localhost:8080
// @BasePath /
func main() {
	db, err := database.NewDB("./todos.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.InitSchema(); err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	todoHandler := &handlers.TodoHandler{DB: db}

	r.HandleFunc("/todos", todoHandler.GetTodos).Methods("GET")
	r.HandleFunc("/todos/{id}", todoHandler.GetTodo).Methods("GET")
	r.HandleFunc("/todos", todoHandler.CreateTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", todoHandler.UpdateTodo).Methods("PUT")
	r.HandleFunc("/todos/{id}", todoHandler.DeleteTodo).Methods("DELETE")

	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4243", "http://localhost:80"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	handler := c.Handler(r)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
