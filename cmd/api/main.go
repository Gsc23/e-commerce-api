package main

import (
	"github.com/Gsc23/e-commerce-api/e-commerce-api/internal/adapter/http"
	"github.com/Gsc23/e-commerce-api/e-commerce-api/pkg/config"
	"github.com/Gsc23/e-commerce-api/e-commerce-api/pkg/database"
	"github.com/Gsc23/e-commerce-api/e-commerce-api/pkg/logger"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		http.HTTPModule(),
		config.ConfigModule(),
		database.DBModule(),
		logger.LoggerModule(),
	)
	app.Run()
}
