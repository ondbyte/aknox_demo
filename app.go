package main

import (
	"fmt"

	"github.com/ondbyte/aknox_demo/auth"
	"github.com/ondbyte/aknox_demo/notes"
	"github.com/ondbyte/aknox_demo/router"
)

func Run(port int) {
	router := router.NewSessionsRouter()
	authRepo := auth.NewRepo()
	authService := auth.NewService(authRepo)
	auth.InitRoutes(router, authService)

	notesRepo := notes.NewRepo()
	notesService := notes.NewService(notesRepo)
	notes.InitRoutes(router, notesService)

	fmt.Sprintf("running on port %v", port)
	router.Run(fmt.Sprintf(":%v", port))
}
