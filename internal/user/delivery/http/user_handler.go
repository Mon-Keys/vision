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

func NewUserHandler(router *router.Router, usecase domain.UserUsecase, su domain.SessionUsecase) {
	handler := &userHandler{
		UserUsecase: usecase,
	}
	router.GET("/api/v1/user", middleware.Cors(
		middleware.ReponseMiddlwareAndLogger(
			middleware.Auth(
				middleware.ReponseMiddlwareAndLogger(handler.GetUser), su))))

	router.GET("/api/v1/users", middleware.Cors(
		middleware.ReponseMiddlwareAndLogger(
			middleware.Auth(
				middleware.ReponseMiddlwareAndLogger(handler.GetUsers), su))))

	router.POST("/api/v1/user/create", middleware.Cors(middleware.ReponseMiddlwareAndLogger(handler.CreateUser)))
	router.OPTIONS("/api/v1/user/create", middleware.Cors(middleware.ReponseMiddlwareAndLogger(handler.CreateUser)))
}

func (h *userHandler) GetUser(ctx *fasthttp.RequestCtx) {
	uid := ctx.UserValue("UID").(*domain.UserSession)
	// aid := ctx.UserValue("AID").(*domain.UserSession)
	if uid == nil {
		ctx.SetStatusCode(fasthttp.StatusForbidden)
		return
	}
	user, account, err := h.UserUsecase.FindUserAccountByID(uid.UserID)
	if err != nil {
		log.Printf(err.Error())
		ctx.SetStatusCode(fasthttp.StatusNotFound)
	}
	fmt.Println(user, account)

	userAccount := new(domain.UserAccount)

	userAccount.Created = user.Created
	userAccount.Name = account.Fullname
	userAccount.UserRoleID = account.RoleID

	ctxBody, err := json.Marshal(userAccount)

	if err != nil {
		log.Printf(err.Error())
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	}

	ctx.SetBody(ctxBody)
}

func (h *userHandler) GetUsers(ctx *fasthttp.RequestCtx) {

	uid := ctx.UserValue("UID").(*domain.UserSession)
	if uid == nil {
		ctx.SetStatusCode(fasthttp.StatusForbidden)
		return
	}
	user, account, err := h.UserUsecase.FindUserAccountByID(uid.UserID)
	if err != nil {
		log.Printf(err.Error())
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	if account.RoleID < 2 {
		ctx.SetStatusCode(fasthttp.StatusForbidden)
		return
	}

	fmt.Println(user, account)

	name := string(ctx.QueryArgs().Peek("name"))
	var users []domain.UserAccountFull

	if name != "" {
		users, err = h.UserUsecase.FindAllByName(name)
	} else {
		users, err = h.UserUsecase.All()
	}

	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	ctxBody, err := json.Marshal(users)

	if err != nil {
		log.Printf(err.Error())
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	}

	ctx.SetBody(ctxBody)
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
