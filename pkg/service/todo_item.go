package service

import (
	"github.com/nikita11044/todo-app"
	"github.com/nikita11044/todo-app/pkg/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemService) Create(userId, listId int, item todo.TodoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		// list does not exist or does not belong to user
		return 0, err
	}

	return s.repo.Create(listId, item)
}

func (s *TodoItemService) GetAll(userId, listId int) ([]todo.TodoItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *TodoItemService) GetById(userId, listId int) (todo.TodoItem, error) {
	return s.repo.GetById(userId, listId)
}

func (s *TodoItemService) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}

func (s *TodoItemService) Update(userId, itemId int, input todo.UpdateItemInput) error {
	return s.repo.Update(userId, itemId, input)
}
