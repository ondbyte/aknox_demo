package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewSessionsRouter() *gin.Engine {
	router := gin.Default()
	cookieStore := cookie.NewStore([]byte("yadunandan_k_s"))
	router.Use(sessions.Sessions("session", cookieStore))
	return router
}
