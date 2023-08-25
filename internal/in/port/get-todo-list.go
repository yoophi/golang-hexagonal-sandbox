package port

import "hexagornal-sandbox/internal/domain/entity"

type GetTodoListUseCase interface {
	GetTodoList() ([]entity.Todo, error)
}
