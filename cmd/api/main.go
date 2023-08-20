package main

import (
	"log"

	"github.com/AbdulwahabNour/movies/config"

	"github.com/AbdulwahabNour/movies/internal/server"
	"github.com/AbdulwahabNour/movies/pkg/db/postgres"
	"github.com/AbdulwahabNour/movies/pkg/db/redis"
	"github.com/AbdulwahabNour/movies/pkg/logger"
)

func main() {

	configFile, err := config.LoadConfig("./config/config-local")

	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	conf, err := config.ParseConfig(configFile)
	if err != nil {
		log.Fatalf("Failed to parse config: %v", err)

	}
	logger := logger.NewApiLogger(conf)

	psql, err := postgres.ConnectSql(conf)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer psql.Client.Close()

	redisClient := redis.NewRedisDB(conf)
	defer redisClient.Client.Close()

	serv := server.NewServer(conf, logger, psql.Client, redisClient.Client)
	err = serv.Run()
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}
