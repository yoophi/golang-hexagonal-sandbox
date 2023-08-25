package web

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func NewRouter(i *do.Injector) (*gin.Engine, error) {
	router := gin.Default()

	return router, nil
}
