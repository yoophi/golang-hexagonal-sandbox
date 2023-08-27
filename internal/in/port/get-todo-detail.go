package port

import (
	"hexagornal-sandbox/internal/domain/entity"
	"hexagornal-sandbox/internal/out/port/repository"
)

type GetTodoDetailUseCase interface {
	GetTodoDetail(repo repository.TodoFetchable, todoId string) (entity.Todo, error)
}
