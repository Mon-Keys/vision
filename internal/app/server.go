package vision

import (
	router "github.com/fasthttp/router"
	"github.com/perlinleo/vision/internal/middleware"
	"github.com/valyala/fasthttp"
)

type Server struct {
	Router *router.Router
	Config *Config
}

func Index(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
}

func NewServer(config *Config) (*Server, error) {
	server := &Server{
		Router: NewRouter(),
		Config: config,
	}

	return server, nil
}
func NewRouter() *router.Router {
	router := router.New()
	router.GET("/", middleware.ReponseMiddlwareAndLogger(Index))

	return router
}
