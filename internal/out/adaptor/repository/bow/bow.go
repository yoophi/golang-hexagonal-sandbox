package bow

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/samber/do"
	"github.com/zippoxer/bow"

	"hexagornal-sandbox/internal/domain/entity"
)

type Repository struct {
}

func NewRepository(injector *do.Injector) (*Repository, error) {
	return &Repository{}, nil
}

func (b Repository) CreateTodo(title string, isCompleted bool) (entity.Todo, error) {
	var (
		err     error
		newTodo TodoModel
		result  entity.Todo
	)
	db, err := bow.Open("test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	newTodo = TodoModel{
		Id:          bow.NewId(),
		Title:       title,
		IsCompleted: isCompleted,
		CreatedAt:   time.Now(),
	}
	err = db.Bucket("todos").Put(newTodo)
	if err != nil {
		log.Fatal(err)
	}

	result = entity.Todo{
		ID:          newTodo.Id.String(),
		Title:       newTodo.Title,
		IsCompleted: newTodo.IsCompleted,
		CreatedAt:   newTodo.CreatedAt,
	}

	return result, nil

}

func (b Repository) GetTodoList() ([]entity.Todo, error) {
	var (
		err    error
		todo   TodoModel
		result []entity.Todo
	)
	db, err := bow.Open("test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	iter := db.Bucket("todos").Iter()
	defer iter.Close()
	for iter.Next(&todo) {
		var tmp TodoModel
		err = db.Bucket("todos").Get(todo.Id, &tmp)
		if err != nil {
			fmt.Printf("====== todo with %+v not found\n", todo.Id)
		} else {
			fmt.Printf("====== todo %+v / %+v\n", todo.Id, tmp)

		}
		result = append(result, entity.Todo{
			ID:          todo.Id.String(),
			Title:       todo.Title,
			IsCompleted: todo.IsCompleted,
			CreatedAt:   todo.CreatedAt,
		})
	}
	if iter.Err() != nil {
		log.Fatal(err)
	}

	return result, nil
}

func (b Repository) GetTodoDetail(todoId string) (entity.Todo, error) {
	var (
		err    error
		todo   TodoModel
		result entity.Todo
	)
	db, err := bow.Open("test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	todoId = "BpoAA20AW2w"

	fmt.Printf("todoId: `%+v`\n", todoId)
	err = db.Bucket("todos").Get(todoId, &todo)
	fmt.Printf("err: %+v\n", err)
	if err != nil {
		return entity.Todo{}, errors.New("todo not found")
	}

	result = entity.Todo{
		ID:          todo.Id.String(),
		Title:       todo.Title,
		IsCompleted: todo.IsCompleted,
		CreatedAt:   todo.CreatedAt,
	}
	return result, nil
}

func (b Repository) DeleteTodo(todoId string) error {
	var (
		err error
	)
	db, err := bow.Open("test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("DeleteTodo", todoId)
	err = db.Bucket("todos").Delete(todoId)
	if err != nil {
		panic(err)
	}

	return nil
}
