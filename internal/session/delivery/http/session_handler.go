package delivery

import (
	"encoding/json"

	"github.com/fasthttp/router"
	"github.com/perlinleo/vision/internal/domain"
	"github.com/perlinleo/vision/internal/middleware"
	"github.com/valyala/fasthttp"
)

type sessionHandler struct {
	SessionUsecase domain.SessionUsecase
}

func NewSessionHandler(router *router.Router, usecase domain.SessionUsecase) {
	handler := &sessionHandler{
		SessionUsecase: usecase,
	}

	router.POST("/api/v1/auth", middleware.Cors(middleware.ReponseMiddlwareAndLogger(handler.Login)))
	router.DELETE("/api/v1/auth", middleware.Cors(middleware.ReponseMiddlwareAndLogger(handler.Logout)))
	router.OPTIONS("/api/v1/auth", middleware.Cors(middleware.ReponseMiddlwareAndLogger(handler.Login)))
}

func (h *sessionHandler) Login(ctx *fasthttp.RequestCtx) {
	loginCredentials := new(domain.LoginCredentials)
	err := json.Unmarshal(ctx.PostBody(), &loginCredentials)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	}
	h.SessionUsecase.Login(*loginCredentials)
}

func (h *sessionHandler) Logout(ctx *fasthttp.RequestCtx) {

}
