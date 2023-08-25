package usecases

import (
	"hexagornal-sandbox/internal/domain/entity"
	"hexagornal-sandbox/internal/in/port"
	"hexagornal-sandbox/internal/out/port/repository"

	"github.com/samber/do"
)

type TodoService struct {
	TodoCreatableRepository repository.TodoCreatable
	TodoFetchableRepository repository.TodoFetchable
}

func NewGetTodoListUseCase(injector *do.Injector) (port.GetTodoListUseCase, error) {
	return &TodoService{
		TodoCreatableRepository: do.MustInvoke[repository.TodoCreatable](injector),
		TodoFetchableRepository: do.MustInvoke[repository.TodoFetchable](injector),
	}, nil
}

func (svc TodoService) GetTodoList() ([]entity.Todo, error) {
	todos, err := svc.TodoFetchableRepository.GetTodo()
	return todos, err
}

func (svc TodoService) CreateTodo(input string) error {
	_, err := svc.TodoCreatableRepository.CreateTodo()
	if err != nil {
		return err
	}
	return nil
}

func NewCreateTodoUseCase(injector *do.Injector) (port.CreateTodoUseCase, error) {
	return &TodoService{
		TodoCreatableRepository: do.MustInvoke[repository.TodoCreatable](injector),
		TodoFetchableRepository: do.MustInvoke[repository.TodoFetchable](injector),
	}, nil
}
