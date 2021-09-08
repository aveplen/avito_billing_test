package main

import (
	"flag"

	"github.com/aveplen/avito_test/internal/config"
	"github.com/aveplen/avito_test/internal/handler"
	"github.com/aveplen/avito_test/internal/repository"
	"github.com/aveplen/avito_test/internal/service"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v4"
	"go.uber.org/zap"
)

var (
	configPath string
)

func main() {
	flag.StringVar(
		&configPath,
		"config-path",
		"config/config.yml",
		"path to configuration file",
	)
	flag.Parse()

	/*=================================================================*/
	/*                            Config                               */
	/*=================================================================*/
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		panic(err)
	}

	/*=================================================================*/
	/*                            Logger                               */
	/*=================================================================*/
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	/*=================================================================*/
	/*                          Database                               */
	/*=================================================================*/
	dbconn, err := repository.NewDBConnectionPool(cfg.Postgres)
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer dbconn.Close()

	/*=================================================================*/
	/*                           Handler                               */
	/*=================================================================*/
	repository := repository.NewRepository(dbconn)
	service := service.NewService(cfg, repository)
	handler := handler.NewHandler(service, logger)

	/*=================================================================*/
	/*                            Server                               */
	/*=================================================================*/
	// gin.SetMode(gin.ReleaseMode)
	g := gin.Default()
	handler.RouteBalance(g)
	g.Run(cfg.BindAddr)
}
