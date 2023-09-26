package notes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ondbyte/aknox_demo/auth"
	"github.com/ondbyte/aknox_demo/response"
)

type Handler struct {
	notesService IService
}

func (h *Handler) getAll(ctx *gin.Context) {
	type Body struct {
		Sid string `json:"sid"`
	}

	body := new(Body)
	err := ctx.ShouldBindJSON(body)
	if err != nil {
		ctx.Error(err)
	}

	user, err := auth.GetUserForSession(ctx, body.Sid)
	if err != nil {
		ctx.Error(fmt.Errorf("user not found"))
	}
	if user == nil {
		ctx.Error(fmt.Errorf("user not found"))
	}
	if response.RespondIfError(ctx, http.StatusTeapot) {
		return
	}
	all, err := h.notesService.getAll(user.Email)
	if err != nil {
		ctx.Error(err)
	}
	if response.RespondIfError(ctx, http.StatusTeapot) {
		return
	}
	response.Respond(ctx, http.StatusOK, all)
}

func (h *Handler) create(ctx *gin.Context) {
	type Body struct {
		Sid  string `json:"sid"`
		Note string `json:"note"`
	}

	body := new(Body)
	ctx.ShouldBindJSON(body)

	user, err := auth.GetUserForSession(ctx, body.Sid)
	if err != nil {
		ctx.Error(fmt.Errorf("user not logged in"))
	}
	if response.RespondIfError(ctx, http.StatusTeapot) {
		return
	}
	all, err := h.notesService.create(user.Email, body.Note)
	if err != nil {
		ctx.Error(err)
	}
	if response.RespondIfError(ctx, http.StatusInternalServerError) {
		return
	}
	response.Respond(ctx, http.StatusOK, all)
}

func (h *Handler) delete(ctx *gin.Context) {
	type Body struct {
		Sid string `json:"sid"`
		Id  string `json:"id"`
	}

	body := new(Body)
	ctx.ShouldBindJSON(body)

	user, err := auth.GetUserForSession(ctx, body.Sid)
	if err != nil || user == nil {
		ctx.Error(fmt.Errorf("user not found"))
	}
	if response.RespondIfError(ctx, http.StatusTeapot) {
		return
	}
	all, err := h.notesService.delete(user.Email, body.Id)
	if err != nil {
		ctx.Error(err)
	}
	if response.RespondIfError(ctx, http.StatusInternalServerError) {
		return
	}
	response.Respond(ctx, http.StatusOK, all)
}

func InitRoutes(router *gin.Engine, notesService IService) error {
	h := &Handler{
		notesService: notesService,
	}
	router.GET("notes", h.getAll)
	router.POST("notes", h.create)
	router.DELETE("notes", h.delete)
	return nil
}
