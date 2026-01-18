package main

import (
	"github.com/Gsc23/e-commerce-api/e-commerce-api/internal/adapter/http"
	"github.com/Gsc23/e-commerce-api/e-commerce-api/pkg/config"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		http.HTTPModule(),
		config.ConfigModule(),
	)
	app.Run()
}
