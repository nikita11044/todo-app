package main

import (
	"log"

	"github.com/nikita11044/todo-app"
	"github.com/nikita11044/todo-app/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running hhtp server: %s", err.Error())
	}
}