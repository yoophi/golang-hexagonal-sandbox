package port

type CreateTodoUseCase interface {
	CreateTodo(input string) error
}
