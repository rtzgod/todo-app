package service

import (
	todo "github.com/rtzgod/todo-app"
	"github.com/rtzgod/todo-app/internal/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}
func (s *TodoListService) Create(userID int, list todo.TodoList) (int, error) {
	return s.repo.Create(userID, list)
}

func (s TodoListService) GetAll(userID int) ([]todo.TodoList, error) {
	return s.repo.GetAll(userID)
}
func (s *TodoListService) GetByID(userId, id int) (todo.TodoList, error) {
	return s.repo.GetByID(userId, id)
}
func (s *TodoListService) Delete(userID, listID int) error {
	return s.repo.Delete(userID, listID)
}
