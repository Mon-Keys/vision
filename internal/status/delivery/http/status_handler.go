package delivery

import (
	"encoding/json"

	"github.com/fasthttp/router"
	"github.com/perlinleo/vision/internal/domain"
	"github.com/perlinleo/vision/internal/middleware"
	"github.com/valyala/fasthttp"
)

type statusHandler struct {
	StatusUsecase domain.StatusUsecase
}

func NewStatusHandler(router *router.Router, usecase domain.StatusUsecase) {
	handler := &statusHandler{
		StatusUsecase: usecase,
	}

	router.GET("/api/v1/status", middleware.ReponseMiddlwareAndLogger(handler.Status))
}

func (h *statusHandler) Status(ctx *fasthttp.RequestCtx) {
	usersAmount, err := h.StatusUsecase.Status()
	ctx.SetStatusCode(fasthttp.StatusOK)

	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	}

	status := &domain.Status{
		UsersAmount: usersAmount.UsersAmount,
	}

	ctxBody, err := json.Marshal(status)

	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	}

	ctx.SetBody(ctxBody)
}
