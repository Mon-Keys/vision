package vision

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type DataBaseConfig struct {
	host string
	port string
	user string
	pass string
	name string
}

type RedisConfig struct {
	addr     string
	password string
	db       int
}

type AppConfig struct {
	SessionKey  string
	Port        string
	DatabaseURL string
}

type Config struct {
	App          *AppConfig
	PSQLDB       *DataBaseConfig
	RedisUser    *RedisConfig
	RedisAccount *RedisConfig
}

func NewConfig() *Config {
	viper.SetConfigFile("../../config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}

	// PSQL
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connection := fmt.Sprintf(
		"host=%s dbname=%s sslmode=disable port=%s password=%s user=%s",
		dbHost, dbName, dbPort, dbPass, dbUser)

	// PORT
	appPort := viper.GetString(`server.address`)

	// Redis-user

	redisUserHost := viper.GetString(`redis-user.host`)
	redisUserPassword := viper.GetString(`redis-user.password`)
	redisUserDB := viper.GetInt(`redis-user.DB`)
	redisUserPort := viper.GetString(`redis-user.port`)
	redisUserAddr := fmt.Sprintf("%s:%s", redisUserHost, redisUserPort)

	// Redis-account

	redisAccountHost := viper.GetString(`redis-account.host`)
	redisAccountPassword := viper.GetString(`redis-account.password`)
	redisAccountDB := viper.GetInt(`redis-account.DB`)
	redisAccountPort := viper.GetString(`redis-account.port`)

	redisAccountAddr := fmt.Sprintf("%s:%s", redisAccountHost, redisAccountPort)

	return &Config{
		RedisUser: &RedisConfig{
			addr:     redisUserAddr,
			password: redisUserPassword,
			db:       redisUserDB,
		},
		RedisAccount: &RedisConfig{
			addr:     redisAccountAddr,
			password: redisAccountPassword,
			db:       redisAccountDB,
		},
		PSQLDB: &DataBaseConfig{
			host: dbHost,
			port: dbPort,
			user: dbUser,
			pass: dbPass,
			name: dbName,
		},
		App: &AppConfig{
			SessionKey:  "dsdsdsds",
			DatabaseURL: connection,
			Port:        appPort,
		},
	}
}
