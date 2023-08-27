package repository

import "hexagornal-sandbox/internal/domain/entity"

type TodoFetchable interface {
	GetTodoList() ([]entity.Todo, error)
	GetTodoDetail(todoId string) (entity.Todo, error)
}

type TodoCreatable interface {
	CreateTodo(title string, isCompleted bool) (entity.Todo, error)
}

type TodoDeletable interface {
	DeleteTodo(id string) error
}
