package main

import (
	"log"

	"github.com/VladislavKuch02/todo-app"
	"github.com/VladislavKuch02/todo-app/package/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error ocoured while running http server: %s", err.Error())
	}
}
