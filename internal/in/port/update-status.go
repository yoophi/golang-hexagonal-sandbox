package port

type UpdateTodo interface {
	UpdateTodo(input string) error
}
