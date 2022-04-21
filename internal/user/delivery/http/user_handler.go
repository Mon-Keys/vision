package delivery

import (
	"encoding/json"
	"fmt"
	"log"

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

	router.POST("/api/v1/user/create", middleware.Cors(middleware.ReponseMiddlwareAndLogger(handler.CreateUser)))
	router.OPTIONS("/api/v1/user/create", middleware.Cors(middleware.ReponseMiddlwareAndLogger(handler.CreateUser)))
}

func (h *userHandler) CreateUser(ctx *fasthttp.RequestCtx) {
	NewUserNoAccount := new(domain.NewUserWithoutAccount)

	err := json.Unmarshal(ctx.PostBody(), &NewUserNoAccount)

	if err != nil {
		log.Printf(err.Error())
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	}

	err = h.UserUsecase.SignUpUser(NewUserNoAccount)

	if err != nil {
		log.Printf(err.Error())
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	}

}

func (h *userHandler) Cors(ctx *fasthttp.RequestCtx) {
	fmt.Println(ctx)
	ctx.SetStatusCode(200)
}
