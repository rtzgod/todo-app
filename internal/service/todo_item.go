package service

import (
	todo "github.com/rtzgod/todo-app"
	"github.com/rtzgod/todo-app/internal/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}
func (s *TodoItemService) Create(userId, listId int, item todo.TodoItem) (int, error) {
	_, err := s.listRepo.GetByID(userId, listId)
	if err == nil {
		return 0, err
	}
	return s.repo.Create(userId, listId, item)
}
