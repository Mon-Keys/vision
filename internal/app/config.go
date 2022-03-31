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
	App     *AppConfig
	PSQLDB  *DataBaseConfig
	RedisDB *RedisConfig
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

	// Redis

	redisHost := viper.GetString(`redis.host`)
	redisPassword := viper.GetString(`redis.password`)
	redisDB := viper.GetInt(`redis.DB`)
	redisPort := viper.GetString(`redis.port`)

	redisAddr := fmt.Sprintf("%s:%s", redisHost, redisPort)

	return &Config{
		RedisDB: &RedisConfig{
			addr:     redisAddr,
			password: redisPassword,
			db:       redisDB,
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
