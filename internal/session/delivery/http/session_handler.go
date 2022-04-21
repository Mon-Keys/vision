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

	router.POST("/api/v1/auth", middleware.Cors(middleware.ReponseMiddlwareAndLogger(handler.Login)))
	router.DELETE("/api/v1/auth", middleware.Cors(middleware.ReponseMiddlwareAndLogger(handler.Logout)))
	router.OPTIONS("/api/v1/auth", middleware.Cors(middleware.ReponseMiddlwareAndLogger(handler.Login)))
}

func (h *sessionHandler) Login(ctx *fasthttp.RequestCtx) {
	loginCredentials := new(domain.LoginCredentials)
	err := json.Unmarshal(ctx.PostBody(), &loginCredentials)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	session, err := h.SessionUsecase.Login(*loginCredentials)

	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusForbidden)
		return
	}

	cookie := cookie.CreateFastHTTPCookie(*session)
	// ctx.Response.Header.Cookie()

	// set := ctx.Response.Header.Cookie(&cookie)

	ctx.Response.Header.Add("Set-Cookie", cookie.String())
	fmt.Println(cookie.String())
	// ctx.Response.Header.Set("Set-Cookie", "value")
	// ctx.Response.Header.Add("Set-Cookie", "<NAME>=<JOPA>")
	// fmt.Printf(string(set))
	// if set {
	// 	return
	// }
}

func (h *sessionHandler) Logout(ctx *fasthttp.RequestCtx) {

}
