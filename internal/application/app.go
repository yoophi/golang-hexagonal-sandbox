package application

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

type Application struct {
	Injector *do.Injector
	Router   *gin.Engine
}

func NewApplication(injector *do.Injector) (*Application, error) {
	router := do.MustInvoke[*gin.Engine](injector)

	return &Application{Injector: injector, Router: router}, nil
}
