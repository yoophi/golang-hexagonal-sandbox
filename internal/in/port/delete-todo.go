package port

type DeleteTodo interface {
	DeleteTodo(input string) error
}
