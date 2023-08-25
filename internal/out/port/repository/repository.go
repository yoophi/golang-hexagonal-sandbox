package repository

import "hexagornal-sandbox/internal/domain/entity"

type TodoFetchable interface {
	GetTodo() ([]entity.Todo, error)
}

type TodoCreatable interface {
	CreateTodo() ([]entity.Todo, error)
}
