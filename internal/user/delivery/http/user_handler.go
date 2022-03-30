package delivery

import (
	"github.com/fasthttp/router"
	"github.com/perlinleo/vision/internal/domain"
	"github.com/perlinleo/vision/internal/middleware"
	"github.com/valyala/fasthttp"
)

type userHandler struct {
	UserUsecase domain.UserUsecase
}

func NewUserHandler(router *router.Router, usecase domain.UserUsecase) {
	handler := &userHandler{
		UserUsecase: usecase,
	}

	router.POST("/api/v1/user/{nickname}/create", middleware.ReponseMiddlwareAndLogger(handler.CreateUser))
}

func (h *userHandler) CreateUser(ctx *fasthttp.RequestCtx) {
	// nickname := ctx.UserValue("nickname").(string)
}
