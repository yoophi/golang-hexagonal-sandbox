package repository

import (
	"time"

	"hexagornal-sandbox/internal/domain/entity"
	"hexagornal-sandbox/internal/out/adaptor/repository/bow"
	"hexagornal-sandbox/internal/out/port/repository"

	"github.com/samber/do"
)

type TodoRepository struct {
}

func (repo TodoRepository) GetTodo() ([]entity.Todo, error) {
	return []entity.Todo{
		{
			Title:       "hello",
			IsCompleted: false,
			CreatedAt:   time.Now(),
		},
	}, nil
}

func (repo TodoRepository) CreateTodo(string, bool) (entity.Todo, error) {
	return entity.Todo{}, nil
}

func NewTodoCreatableRepository(injector *do.Injector) (repository.TodoCreatable, error) {
	return bow.NewRepository(injector)
}

func NewTodoDeletableRepository(injector *do.Injector) (repository.TodoDeletable, error) {
	return bow.NewRepository(injector)
}

func NewTodoFetchableRepository(injector *do.Injector) (repository.TodoFetchable, error) {
	return bow.NewRepository(injector)
}
