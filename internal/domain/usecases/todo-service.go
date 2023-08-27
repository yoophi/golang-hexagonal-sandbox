package usecases

import (
	"errors"
	"fmt"

	"hexagornal-sandbox/internal/domain/entity"
	"hexagornal-sandbox/internal/in/port"
	"hexagornal-sandbox/internal/in/port/req"
	"hexagornal-sandbox/internal/out/port/repository"

	"github.com/samber/do"
)

type TodoService struct {
}

func (svc *TodoService) GetTodoDetail(repo repository.TodoFetchable, todoId string) (entity.Todo, error) {
	todo, err := repo.GetTodoDetail(todoId)
	fmt.Printf("%+v\n", todo)
	if err != nil {
		return entity.Todo{}, errors.New("todo not found")
	}

	return todo, nil
}

func (svc *TodoService) GetTodoList(repo repository.TodoFetchable) ([]entity.Todo, error) {
	todos, err := repo.GetTodoList()
	return todos, err
}

func (svc *TodoService) CreateTodo(repo repository.TodoCreatable, input req.CreateTodoRequest) error {
	_, err := repo.CreateTodo(input.Title, input.IsCompleted)
	if err != nil {
		return errors.New("create todo failed")
	}
	return nil
}

func (svc *TodoService) DeleteTodo(repo repository.TodoDeletable, todoId string) error {
	var (
		err error
	)

	fmt.Print("delete todo", todoId)
	err = repo.DeleteTodo(todoId)
	if err != nil {
		return errors.New("delete todo failed")
	}

	return nil
}

func NewGetTodoListUseCase(injector *do.Injector) (port.GetTodoListUseCase, error) {
	return &TodoService{}, nil
}

func NewGetTodoDetailUseCase(injector *do.Injector) (port.GetTodoDetailUseCase, error) {
	return &TodoService{}, nil
}

func NewCreateTodoUseCase(injector *do.Injector) (port.CreateTodoUseCase, error) {
	return &TodoService{}, nil
}

func NewDeleteTodoUseCase(injector *do.Injector) (port.DeleteTodoUseCase, error) {
	return &TodoService{}, nil
}
