package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ondbyte/aknox_demo/response"
)

type Handler struct {
	authService IService
}

func (h *Handler) signUp(ctx *gin.Context) {
	type Body struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	body := &Body{}
	err := ctx.Bind(body)
	if err != nil {
		ctx.Error(err)
	}
	//currently we are not validating any email
	if body.Name == "" {
		ctx.Error(fmt.Errorf("name is required"))
	}
	if body.Email == "" {
		ctx.Error(fmt.Errorf("email is required"))
	}
	if body.Password == "" {
		ctx.Error(fmt.Errorf("password is required"))
	}
	if response.RespondIfError(ctx, http.StatusTeapot) {
		return
	}
	_, err = h.authService.SignUp(body.Name, body.Email, body.Password)
	if err != nil {
		ctx.Error(err)
	}
	response.Respond(ctx, http.StatusOK, "success")
}

func (h *Handler) login(ctx *gin.Context) {
	type Body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	body := &Body{}
	err := ctx.Bind(body)
	if err != nil {
		ctx.Error(err)
	}
	if body.Email == "" {
		ctx.Error(fmt.Errorf("email is required"))
	}
	if body.Password == "" {
		ctx.Error(fmt.Errorf("password is required"))
	}
	user, err := h.authService.Login(body.Email, body.Password)
	if err != nil {
		ctx.Error(err)
	}
	if response.RespondIfError(ctx, http.StatusTeapot) {
		return
	}
	//login is successful generate session
	sid := NewSessionId()
	SetUserForSession(ctx, sid, *user)
	response.Respond(ctx, http.StatusOK, map[string]string{"sid": sid})
}

func InitRoutes(router *gin.Engine, authService IService) error {
	h := &Handler{
		authService: authService,
	}
	router.POST("/signup", h.signUp)
	router.POST("/login", h.login)
	return nil
}
