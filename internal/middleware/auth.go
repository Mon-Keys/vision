package middleware

import (
	"fmt"

	"github.com/perlinleo/vision/internal/domain"
	"github.com/valyala/fasthttp"
)

func Auth(next fasthttp.RequestHandler, sessionRepository domain.SessionUsecase) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		accountCookie := ctx.Request.Header.Cookie("UID")
		account, _ := sessionRepository.GetUserSessionByCookie(string(accountCookie))
		// if err != nil {
		// 	log.Printf(err.Error())
		// 	return
		// }

		sessionCookie := ctx.Request.Header.Cookie("AID")
		session, _ := sessionRepository.GetUserSessionByCookie(string(sessionCookie))
		// if err != nil {
		// 	log.Printf(err.Error())
		// 	return
		// }

		ctx.SetUserValue("UID", account)
		ctx.SetUserValue("AID", session)
		fmt.Println(account)
		fmt.Println(session)
		next(ctx)
	}
}
