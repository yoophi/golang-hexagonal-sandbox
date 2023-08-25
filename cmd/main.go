package main

import (
	"errors"
	"fmt"
	"hexagornal-sandbox/internal/application"
	"hexagornal-sandbox/internal/domain/usecases"
	"hexagornal-sandbox/internal/out/adaptor/repository"
	"log"
	"net/http"

	"hexagornal-sandbox/internal/in/adaptor/web"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func main() {
	injector := do.New()
	do.Provide(injector, application.NewApplication)
	do.Provide(injector, web.NewRouter)
	do.Provide(injector, web.NewTodoController)
	do.Provide(injector, usecases.NewGetTodoListUseCase)
	do.Provide(injector, usecases.NewCreateTodoUseCase)
	do.Provide(injector, repository.NewTodoCreatableRepository)
	do.Provide(injector, repository.NewTodoFetchableRepository)

	injectedServices := injector.ListProvidedServices()
	for idx, s := range injectedServices {
		fmt.Printf("%d - %#v\n", idx, s)
	}

	app := do.MustInvoke[*application.Application](injector)
	do.MustInvoke[*web.TodoController](injector)
	fmt.Printf("%#v\n", app)

	router := do.MustInvoke[*gin.Engine](injector)
	srv := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("listen: %#v\n", err)
	}
}
