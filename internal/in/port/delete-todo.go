package port

import "hexagornal-sandbox/internal/out/port/repository"

type DeleteTodoUseCase interface {
	DeleteTodo(repo repository.TodoDeletable, todoId string) error
}
