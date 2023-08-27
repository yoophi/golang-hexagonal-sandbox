package web

import (
	"net/http"

	"hexagornal-sandbox/internal/application"
	"hexagornal-sandbox/internal/in/port"
	"hexagornal-sandbox/internal/in/port/req"
	"hexagornal-sandbox/internal/out/port/repository"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

type TodoController struct {
	app application.Application
}

func (ctrl TodoController) GetTodoList(c *gin.Context) {
	var (
		injector = ctrl.app.Injector
		repo     = do.MustInvoke[repository.TodoFetchable](injector)
		usecase  = do.MustInvoke[port.GetTodoListUseCase](injector)
	)
	todos, err := usecase.GetTodoList(repo)
	if err != nil {
		return
	}

	c.JSON(200, gin.H{"message": todos})
}

func (ctrl TodoController) GetTodoDetail(c *gin.Context) {
	var (
		injector = ctrl.app.Injector
		repo     = do.MustInvoke[repository.TodoFetchable](injector)
		usecase  = do.MustInvoke[port.GetTodoDetailUseCase](injector)
		todoId   string
	)
	todoId = c.Param("todo-id")

	todo, err := usecase.GetTodoDetail(repo, todoId)
	if err != nil {
		return
	}

	c.JSON(200, gin.H{"message": todo})
}

func (ctrl TodoController) RegisterTodo(c *gin.Context) {
	type (
		RegisterTodoRequest struct {
			Title       string `json:"title"`
			IsCompleted bool   `json:"is_completed"`
		}
	)
	var (
		injector = ctrl.app.Injector
		repo     = do.MustInvoke[repository.TodoCreatable](injector)
		usecase  = do.MustInvoke[port.CreateTodoUseCase](injector)
		err      error
		payload  RegisterTodoRequest
	)

	if err = c.ShouldBindJSON(&payload); err != nil {
		panic(err)
	}
	request := req.CreateTodoRequest{
		Title:       payload.Title,
		IsCompleted: payload.IsCompleted,
	}
	if err = usecase.CreateTodo(repo, request); err != nil {
		return
	}
	c.JSON(200, gin.H{})
}

func (ctrl TodoController) UpdateTodo() {

}

func (ctrl TodoController) SetTodoState() {

}

func (ctrl TodoController) RemoveTodo(c *gin.Context) {
	var (
		injector = ctrl.app.Injector
		repo     = do.MustInvoke[repository.TodoDeletable](injector)
		usecase  = do.MustInvoke[port.DeleteTodoUseCase](injector)
		err      error
		todoId   string
	)
	usecase = do.MustInvoke[port.DeleteTodoUseCase](injector)
	repo = do.MustInvoke[repository.TodoDeletable](injector)
	todoId = c.Param("todo-id")
	err = usecase.DeleteTodo(repo, todoId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (ctrl TodoController) RegisterRoutes(route *gin.RouterGroup) error {
	route.GET("", ctrl.GetTodoList)
	route.POST("", ctrl.RegisterTodo)
	route.GET(":todo-id", ctrl.GetTodoDetail)
	route.DELETE(":todo-id", ctrl.RemoveTodo)
	return nil
}

func NewTodoController(injector *do.Injector) (*TodoController, error) {
	app := do.MustInvoke[*application.Application](injector)
	router := app.Router
	controller := &TodoController{
		app: *app,
	}
	err := controller.RegisterRoutes(router.Group("/api/v1/todos"))
	if err != nil {
		return nil, err
	}

	return controller, nil
}
