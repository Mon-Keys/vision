package vision

import (
	"log"
	"time"

	router "github.com/fasthttp/router"
	"github.com/patrickmn/go-cache"

	// forum_psql "github.com/perlinleo/technopark-mail.ru-forum-database/internal/app/forum/repository"
	// thread_psql "github.com/perlinleo/technopark-mail.ru-forum-database/internal/app/thread/repository"

	// forum_usecase "github.com/perlinleo/technopark-mail.ru-forum-database/internal/app/forum/usecase"
	// thread_usecase "github.com/perlinleo/technopark-mail.ru-forum-database/internal/app/thread/usecase"

	user_http "github.com/perlinleo/vision/internal/user/delivery/http"
	user_psql "github.com/perlinleo/vision/internal/user/repository/psql"

	user_usecase "github.com/perlinleo/vision/internal/user/usecase"

	status_http "github.com/perlinleo/vision/internal/status/delivery/http"
	status_psql "github.com/perlinleo/vision/internal/status/repository/psql"
	status_usecase "github.com/perlinleo/vision/internal/status/usecase"

	"github.com/valyala/fasthttp"
	// forum_http "github.com/perlinleo/technopark-mail.ru-forum-database/internal/app/forum/delivery"
	// thread_http "github.com/perlinleo/technopark-mail.ru-forum-database/internal/app/thread/delivery"
)

func Start() error {
	config := NewConfig()

	_, err := NewServer(config)
	if err != nil {
		return err
	}
	PSQLConnPool, err := NewPostgreSQLDataBase(config.App.DatabaseURL)
	if err != nil {
		return err
	}

	router := router.New()

	userCache := cache.New(5*time.Minute, 10*time.Minute)

	statusCache := cache.New(5*time.Minute, 10*time.Minute)

	userRepository := user_psql.NewUserPSQLRepository(PSQLConnPool, userCache)
	userUsecase := user_usecase.NewUserUsecase(userRepository)
	user_http.NewUserHandler(router, userUsecase)

	statusRepository := status_psql.NewStatusPSQLRepository(PSQLConnPool, statusCache)
	statusUsecase := status_usecase.NewStatusUsecase(statusRepository)
	status_http.NewStatusHandler(router, statusUsecase)

	log.Printf("STARTING SERVICE ON PORT %s\n", config.App.Port)

	err = fasthttp.ListenAndServe(config.App.Port, router.Handler)
	if err != nil {
		return err
	}

	return nil
}
