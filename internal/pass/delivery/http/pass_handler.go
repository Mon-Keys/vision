package delivery

import (
	"encoding/json"
	"log"

	"github.com/fasthttp/router"
	"github.com/perlinleo/vision/internal/domain"
	"github.com/perlinleo/vision/internal/middleware"
	"github.com/valyala/fasthttp"
)

type passHandler struct {
	PassUsecase domain.PassUsecase
	UserUsecase domain.UserUsecase
}

func NewPassesHandler(router *router.Router, usecase domain.PassUsecase, user domain.UserUsecase, su domain.SessionUsecase) {
	handler := &passHandler{
		PassUsecase: usecase,
		UserUsecase: user,
	}

	router.GET("/api/v1/passes", middleware.Cors(
		middleware.ReponseMiddlwareAndLogger(
			middleware.Auth(
				middleware.ReponseMiddlwareAndLogger(handler.Passes), su))))

	router.GET("/api/v1/check/{data}", middleware.ReponseMiddlwareAndLogger(handler.Check))

}

func (h *passHandler) Passes(ctx *fasthttp.RequestCtx) {
	aid := ctx.UserValue("AID").(*domain.UserSession)
	passes, err := h.PassUsecase.GetUserPasses(aid.UserID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
	}

	ctxBody, err := json.Marshal(passes)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	}

	ctx.SetBody(ctxBody)
}

func (h *passHandler) Check(ctx *fasthttp.RequestCtx) {
	data := ctx.UserValue("data").(string)
	checkResult, err := h.PassUsecase.CheckPass(data)
	if err != nil {
		log.Printf(err.Error())
	}
	ctxBody, err := json.Marshal(checkResult)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	}
	ctx.SetBody(ctxBody)
}
