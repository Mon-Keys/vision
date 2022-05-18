package middleware

import (
	"fmt"
	"strings"

	"github.com/valyala/fasthttp"
)

var allowedOrigins = map[string]struct{}{
	"http://127.0.0.1":            {},
	"http://127.0.0.1:8000":       {},
	"http://localhost":            {},
	"http://localhost:8080":       {},
	"http://vision.leonidperl.in": {},
	"http://192.168.1.16:8080":    {},

	"https://localhost:8080":    {},
	"https://127.0.0.1:8000":    {},
	"https://192.168.1.16:8080": {},
	"https://127.0.0.1":         {},
	"https://localhost":         {},
	"https://ijia.me":           {},
}

func Cors(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		fmt.Println("CORS")
		var sb strings.Builder
		sb.WriteString("Accept,")
		sb.WriteString("Content-Type,")
		sb.WriteString("Content-Length,")
		sb.WriteString("Accept-Encoding,")
		sb.WriteString("X-CSRF-Token,")
		sb.WriteString("Authorization,")
		sb.WriteString("Allow-Credentials,")
		sb.WriteString("Set-Cookie,")
		sb.WriteString("Access-Control-Allow-Credentials,")
		sb.WriteString("Access-Control-Allow-Origin")
		ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
		ctx.Response.Header.Set("Access-Control-Allow-Methods", "POST, GET, DELETE, PUT, OPTIONS")
		ctx.Response.Header.Set("Access-Control-Allow-Headers", sb.String())
		ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
		ctx.Response.Header.SetBytesV("Access-Control-Allow-Origin", ctx.Request.Header.Peek("Origin"))
		ctx.SetStatusCode(fasthttp.StatusOK)
		if string(ctx.Method()) == fasthttp.MethodOptions {
			return
		}

		next(ctx)
	}
}
