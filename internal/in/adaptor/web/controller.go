package web

import (
	"fmt"
	"hexagornal-sandbox/internal/application"
	"hexagornal-sandbox/internal/in/port"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

type TodoController struct {
	CreateTodoUseCase port.CreateTodoUseCase
	app               application.Application
}

func (ctrl TodoController) GetTodoList(c *gin.Context) {
	usecase := do.MustInvoke[port.GetTodoListUseCase](ctrl.app.Injector)
	todos, err := usecase.GetTodoList()
	if err != nil {
		return
	}
	fmt.Println(todos)
	c.JSON(200, gin.H{"message": todos})
}

func (ctrl TodoController) RegisterTodo(c *gin.Context) {
	usecase := ctrl.CreateTodoUseCase
	err := usecase.CreateTodo("test")
	if err != nil {
		return
	}
	c.JSON(200, gin.H{})
}

func (ctrl TodoController) UpdateTodo() {

}

func (ctrl TodoController) SetTodoState() {

}

func (ctrl TodoController) RemoveTodo() {

}

func (ctrl TodoController) RegisterRoutes(route *gin.RouterGroup) error {
	route.GET("", ctrl.GetTodoList)
	route.POST("", ctrl.RegisterTodo)
	return nil
}

func NewTodoController(injector *do.Injector) (*TodoController, error) {
	app := do.MustInvoke[*application.Application](injector)
	router := do.MustInvoke[*gin.Engine](injector)
	controller := &TodoController{
		app:               *app,
		CreateTodoUseCase: do.MustInvoke[port.CreateTodoUseCase](injector),
	}
	err := controller.RegisterRoutes(router.Group("/api/v1/todos"))
	if err != nil {
		return nil, err
	}

	return controller, nil
}
