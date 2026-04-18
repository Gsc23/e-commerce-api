package main

import (
	"github.com/Gsc23/e-commerce-api/internal/adapter/http"
	"github.com/Gsc23/e-commerce-api/pkg/config"
	"github.com/Gsc23/e-commerce-api/pkg/database"
	"github.com/Gsc23/e-commerce-api/pkg/logger"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		config.ConfigModule(),
		logger.LoggerModule(),
		database.DBModule(),
		http.HTTPModule(),
	)
	app.Run()
}
