package port

import (
	"hexagornal-sandbox/internal/in/port/req"
	"hexagornal-sandbox/internal/out/port/repository"
)

type CreateTodoUseCase interface {
	CreateTodo(repo repository.TodoCreatable, input req.CreateTodoRequest) error
}
