package repository

import (
	"fmt"
	"hexagornal-sandbox/internal/domain/entity"
	"hexagornal-sandbox/internal/out/port/repository"
	"time"

	"github.com/samber/do"
)

type TodoRepository struct {
}

func (repo TodoRepository) GetTodo() ([]entity.Todo, error) {
	return []entity.Todo{
		entity.Todo{
			Title:       "hello",
			IsCompleted: false,
			CreatedAt:   time.Now(),
		},
	}, nil
}

func (repo TodoRepository) CreateTodo() ([]entity.Todo, error) {
	//TODO implement me
	//panic("implement me")
	fmt.Println("Repository: TodoCreatableRepository")
	return nil, nil
}

func NewTodoCreatableRepository(injector *do.Injector) (repository.TodoCreatable, error) {
	return &TodoRepository{}, nil
}

func NewTodoFetchableRepository(injector *do.Injector) (repository.TodoFetchable, error) {
	return &TodoRepository{}, nil
}
