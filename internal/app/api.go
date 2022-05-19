package vision

import (
	"log"
	"time"

	router "github.com/fasthttp/router"
	"github.com/patrickmn/go-cache"

	session_http "github.com/perlinleo/vision/internal/session/delivery/http"
	session_redis "github.com/perlinleo/vision/internal/session/repository/redis"
	session_usecase "github.com/perlinleo/vision/internal/session/usecase"

	user_http "github.com/perlinleo/vision/internal/user/delivery/http"
	user_psql "github.com/perlinleo/vision/internal/user/repository/psql"
	user_usecase "github.com/perlinleo/vision/internal/user/usecase"

	declaration_http "github.com/perlinleo/vision/internal/declaration/delivery/http"
	declaration_psql "github.com/perlinleo/vision/internal/declaration/repository/psql"
	declaration_usecase "github.com/perlinleo/vision/internal/declaration/usecase"

	account_psql "github.com/perlinleo/vision/internal/account/repository/psql"
	pass_http "github.com/perlinleo/vision/internal/pass/delivery/http"
	pass_psql "github.com/perlinleo/vision/internal/pass/repository/psql"
	pass_usecase "github.com/perlinleo/vision/internal/pass/usecase"

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
	log.Printf("PSQL Database connection success on %s", config.App.DatabaseURL)

	RedisClientUser, err := NewRedisDataBase(
		config.RedisUser.addr,
		config.RedisUser.password,
		config.RedisUser.db)

	RedisClientAccount, err := NewRedisDataBase(
		config.RedisAccount.addr,
		config.RedisAccount.password,
		config.RedisAccount.db)

	if err != nil {
		return err
	}

	log.Printf("Redis Database connection success on %s", config.App.DatabaseURL)

	router := router.New()

	// pass
	passCache := cache.New(5*time.Minute, 10*time.Minute)
	passRepository := pass_psql.NewPassPSQLRepository(PSQLConnPool, passCache)
	passUsecase := pass_usecase.NewPassUsecase(passRepository)

	//declaration
	declarationCache := cache.New(5*time.Minute, 10*time.Minute)
	declarationRepository := declaration_psql.NewDeclarationPSQLRepository(PSQLConnPool, declarationCache)

	// account
	accountCache := cache.New(5*time.Minute, 10*time.Minute)
	accountRepository := account_psql.NewAccountPSQLRepository(PSQLConnPool, accountCache)
	declarationUsecase := declaration_usecase.NewDeclarationUsecase(declarationRepository, passRepository, accountRepository)

	// usERR
	userCache := cache.New(5*time.Minute, 10*time.Minute)
	userRepository := user_psql.NewUserPSQLRepository(PSQLConnPool, userCache)
	userUsecase := user_usecase.NewUserUsecase(userRepository, accountRepository)

	//auth
	authCache := cache.New(5*time.Minute, 10*time.Minute)
	authRepository := session_redis.NewSessionRedisRepository(&RedisClientUser, &RedisClientAccount, authCache)
	authUsecase := session_usecase.NewSessionUsecase(authRepository, userRepository, accountRepository)

	//declaration
	declaration_http.NewDeclarationHandler(router, declarationUsecase, userUsecase, authUsecase)
	session_http.NewSessionHandler(router, authUsecase)
	user_http.NewUserHandler(router, userUsecase, authUsecase)
	pass_http.NewPassesHandler(router, passUsecase, userUsecase, authUsecase)

	// status
	statusCache := cache.New(5*time.Minute, 10*time.Minute)
	statusRepository := status_psql.NewStatusPSQLRepository(PSQLConnPool, statusCache)
	statusUsecase := status_usecase.NewStatusUsecase(statusRepository)
	status_http.NewStatusHandler(router, statusUsecase)

	log.Printf("STARTING SERVICE ON PORT %s\n", config.App.Port)

	err = fasthttp.ListenAndServe(config.App.Port, router.Handler)
	// err = fasthttp.ListenAndServeTLS(config.App.Port, "cert.pem", "key.pem", router.Handler)
	if err != nil {
		return err
	}

	return nil
}
