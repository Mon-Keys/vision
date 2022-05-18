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

type declarationHandler struct {
	DeclarationUsecase domain.DeclarationUsecase
	UserUsecase        domain.UserUsecase
}

func NewDeclarationHandler(router *router.Router, usecase domain.DeclarationUsecase, user domain.UserUsecase, su domain.SessionUsecase) {
	handler := &declarationHandler{
		DeclarationUsecase: usecase,
		UserUsecase:        user,
	}

	router.GET("/api/v1/declarations", middleware.Cors(
		middleware.ReponseMiddlwareAndLogger(
			middleware.Auth(
				middleware.ReponseMiddlwareAndLogger(handler.Declarations), su))))

	router.POST("/api/v1/asktime", middleware.Cors(
		middleware.ReponseMiddlwareAndLogger(
			middleware.Auth(
				middleware.ReponseMiddlwareAndLogger(handler.AskTime), su))))

	router.POST("/api/v1/askpass", middleware.Cors(
		middleware.ReponseMiddlwareAndLogger(
			middleware.Auth(
				middleware.ReponseMiddlwareAndLogger(handler.AskPass), su))))

	router.POST("/api/v1/askrole", middleware.Cors(
		middleware.ReponseMiddlwareAndLogger(
			middleware.Auth(
				middleware.ReponseMiddlwareAndLogger(handler.AskRole), su))))

	router.POST("/api/v1/declarations/deny", middleware.Cors(
		middleware.ReponseMiddlwareAndLogger(
			middleware.Auth(
				middleware.ReponseMiddlwareAndLogger(handler.Deny), su))))
	router.POST("/api/v1/declarations/accept", middleware.Cors(
		middleware.ReponseMiddlwareAndLogger(
			middleware.Auth(
				middleware.ReponseMiddlwareAndLogger(handler.Accept), su))))
}

func (h *declarationHandler) Accept(ctx *fasthttp.RequestCtx) {
	decToAccept := new(domain.DeclarationCommon)

	err := json.Unmarshal(ctx.PostBody(), &decToAccept)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	}
	uid := ctx.UserValue("UID").(*domain.UserSession)
	aid := ctx.UserValue("AID").(*domain.UserSession)
	fmt.Println(decToAccept, uid.UserID, aid.UserID)

	_, account, err := h.UserUsecase.FindUserAccountByID(uid.UserID)

	if err != nil {
		log.Printf(err.Error())
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}
	if account.RoleID < 3 {
		ctx.SetStatusCode(fasthttp.StatusForbidden)
		return
	}

	err = h.DeclarationUsecase.AcceptDeclaration(*decToAccept)
	if err != nil {
		log.Printf(err.Error())
	}
}

func (h *declarationHandler) Deny(ctx *fasthttp.RequestCtx) {
	decToDeny := new(domain.DeclarationCommon)

	err := json.Unmarshal(ctx.PostBody(), &decToDeny)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	}
	uid := ctx.UserValue("UID").(*domain.UserSession)
	aid := ctx.UserValue("AID").(*domain.UserSession)

	_, account, err := h.UserUsecase.FindUserAccountByID(uid.UserID)

	if err != nil {
		log.Printf(err.Error())
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}
	if account.RoleID < 3 {
		ctx.SetStatusCode(fasthttp.StatusForbidden)
		return
	}
	err = h.DeclarationUsecase.DenyDeclaration(*decToDeny)
	if err != nil {
		log.Printf(err.Error())
	}
	fmt.Println(decToDeny, uid.UserID, aid.UserID)

}

func (h *declarationHandler) Declarations(ctx *fasthttp.RequestCtx) {
	uid := ctx.UserValue("UID").(*domain.UserSession)
	if uid == nil {
		ctx.SetStatusCode(fasthttp.StatusForbidden)
		return
	}
	_, account, err := h.UserUsecase.FindUserAccountByID(uid.UserID)
	if err != nil {
		log.Printf(err.Error())
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}
	var declarations []domain.DeclarationCommon
	if account.RoleID < 2 {
		declarations, err = h.DeclarationUsecase.AllDeclarationsByID(account.ID)
	} else {
		declarations, err = h.DeclarationUsecase.AllDeclarations()
	}
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	}
	ctxBody, err := json.Marshal(declarations)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	}
	ctx.SetBody(ctxBody)
}

func (h *declarationHandler) AskTime(ctx *fasthttp.RequestCtx) {

}

func (h *declarationHandler) AskPass(ctx *fasthttp.RequestCtx) {
	askPass := new(domain.AskPass)
	aid := ctx.UserValue("AID").(*domain.UserSession)

	err := json.Unmarshal(ctx.PostBody(), &askPass)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	}

	fmt.Println(askPass)

	err = h.DeclarationUsecase.CreatePassDeclaration(*askPass, aid.UserID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusConflict)
	}
}

func (h *declarationHandler) AskRole(ctx *fasthttp.RequestCtx) {
	askRole := new(domain.AskRole)

	err := json.Unmarshal(ctx.PostBody(), &askRole)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	}
	uid := ctx.UserValue("UID").(*domain.UserSession)
	aid := ctx.UserValue("AID").(*domain.UserSession)
	fmt.Println(askRole, uid.UserID, aid.UserID)

	askRoleDeclaration := new(domain.AskRoleDeclaration)
	askRoleDeclaration.CreatorID = aid.UserID
	askRoleDeclaration.Comment = askRole.Comment
	askRoleDeclaration.RoleID = askRole.RoleID

	err = h.DeclarationUsecase.CreateRoleDeclaration(*askRoleDeclaration)

	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusForbidden)
	}
}
