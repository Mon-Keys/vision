package delivery

import (
	"encoding/json"
	"fmt"

	"github.com/fasthttp/router"
	"github.com/perlinleo/vision/internal/domain"
	"github.com/perlinleo/vision/internal/middleware"
	cookie "github.com/perlinleo/vision/internal/pkg/cookie_generator"
	"github.com/valyala/fasthttp"
)

type sessionHandler struct {
	SessionUsecase domain.SessionUsecase
}

func NewSessionHandler(router *router.Router, usecase domain.SessionUsecase) {
	handler := &sessionHandler{
		SessionUsecase: usecase,
	}

	router.POST("/api/v1/auth", middleware.Cors(middleware.Auth(middleware.ReponseMiddlwareAndLogger(handler.Login), handler.SessionUsecase)))
	router.DELETE("/api/v1/auth", middleware.Cors(middleware.Auth(middleware.ReponseMiddlwareAndLogger(handler.Logout), handler.SessionUsecase)))
	router.OPTIONS("/api/v1/auth", middleware.Cors(middleware.ReponseMiddlwareAndLogger(handler.Login)))
}

func (h *sessionHandler) Login(ctx *fasthttp.RequestCtx) {
	loginCredentials := new(domain.LoginCredentials)
	err := json.Unmarshal(ctx.PostBody(), &loginCredentials)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	userSession, accountSession, err := h.SessionUsecase.Login(*loginCredentials)

	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusForbidden)
		return
	}

	cookieUser := cookie.CreateFastHTTPCookie(userSession.Cookie, "UID")

	cookieAccount := cookie.CreateFastHTTPCookie(accountSession.Cookie, "AID")

	ctx.Response.Header.Add("Set-Cookie", cookieUser.String())
	ctx.Response.Header.Add("Set-Cookie", cookieAccount.String())
	// fmt.Println(cookie.String())
}

func (h *sessionHandler) Logout(ctx *fasthttp.RequestCtx) {
	uid := ctx.UserValue("UID").(*domain.UserSession)
	aid := ctx.UserValue("AID").(*domain.UserSession)
	fmt.Println(uid.Cookie, aid.Cookie)
	err := h.SessionUsecase.Logout(aid.Cookie, uid.Cookie)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
	}
}
