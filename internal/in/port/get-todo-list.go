package port

import (
	"hexagornal-sandbox/internal/domain/entity"
	"hexagornal-sandbox/internal/out/port/repository"
)

type GetTodoListUseCase interface {
	GetTodoList(repo repository.TodoFetchable) ([]entity.Todo, error)
}
