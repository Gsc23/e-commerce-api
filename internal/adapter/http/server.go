package http

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/Gsc23/e-commerce-api/e-commerce-api/pkg/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Engine interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type Server interface {
	Engine() Engine
	Run(context.Context) error
	Stop(context.Context) error
}

type server struct {
	engine     Engine
	httpSrv    *http.Server
	shutdowner fx.Shutdowner
}

func (s *server) Run(ctx context.Context) error {
	ln, err := net.Listen("tcp", s.httpSrv.Addr)
	if err != nil {
		return err
	}

	fmt.Println("Starting HTTP server at", s.httpSrv.Addr)

	go s.httpSrv.Serve(ln)
	return nil
}

func (s *server) Stop(ctx context.Context) error {
	return s.httpSrv.Shutdown(ctx)
}

func (s *server) Engine() Engine {
	return s.engine
}

func newServer(cfg config.Config, shutdowner fx.Shutdowner) Server {
	var mode string
	switch cfg.Server.Env {
	case "production":
		mode = gin.ReleaseMode
	case "testing":
		mode = gin.TestMode
	default:
		mode = gin.DebugMode
	}

	gin.SetMode(mode)
	engine := gin.Default()
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{cfg.Server.Host, "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	s := &server{
		engine: engine,
		httpSrv: &http.Server{
			Addr:    fmt.Sprintf("%d", cfg.Server.Port),
			Handler: engine.Handler(),
		},
		shutdowner: shutdowner,
	}

	// missing route setups would go here

	return s
}
