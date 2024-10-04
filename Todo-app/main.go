package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"todo/Database"
	CreateTodo "todo/Handler"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the app - Home Page of the todo "))
}

func main() {

	fmt.Println("Welcome to the app - ToDo")

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/v1", func(r chi.Router) {

		r.Get("/home", Home)
		r.Get("/allNotes", CreateTodo.GetAll)
		r.Get("/Notes/{id}", CreateTodo.GetById)
		r.Post("/Note", CreateTodo.CreateNote)
		r.Put("/updateNote", CreateTodo.UpdateTodo)
		r.Delete("/Note/{id}", CreateTodo.DeleteNote)

		//user
		r.Post("/createuser", CreateTodo.CreateUser)

	})

	http.ListenAndServe(":8000", r)

	err := Database.DBConnection
	if err != nil {
		log.Printf("error connecting DaB %v", err)
	}
	defer Database.CloseDatabase()

}
