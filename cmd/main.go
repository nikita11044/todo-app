package main

import (
	"log"

	"github.com/nikita11044/todo-app"
	"github.com/nikita11044/todo-app/pkg/handler"
	"github.com/nikita11044/todo-app/pkg/repository"
	"github.com/nikita11044/todo-app/pkg/service"
)

func main() {
	repo := repository.NewRepository()
	service := service.NewService(repo)
	handler := handler.NewHandler(service)
	
	srv := new(todo.Server)
	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while running hhtp server: %s", err.Error())
	}
}