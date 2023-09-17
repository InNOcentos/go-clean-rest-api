package app

import (
	"context"
	"log"
	"net/http"
	"time"

	userhttp "github.com/InNOcentos/go-clean-rest-api/internal/user/delivery/http"
	userRe "github.com/InNOcentos/go-clean-rest-api/internal/user/repository/postgres"
	userUc "github.com/InNOcentos/go-clean-rest-api/internal/user/usecase"
	"github.com/gin-gonic/gin"

	"github.com/spf13/viper"

	"github.com/InNOcentos/go-clean-rest-api/pkg/database"
)

type App struct {
	httpServer *http.Server
}

func (a *App) Run() {
	if err := a.httpServer.ListenAndServe(); err != nil {
		log.Fatalf("Failed to listen and serve: %+v", err)
	}
}

func NewApp(port string) *App {
	db := InitDb()

	userRepo := userRe.NewRepository(db)
	userUseCase := userUc.NewUseCase(userRepo)

	router := gin.Default()
	router.Use(
		gin.Recovery(),
	)
	userhttp.RegisterHTTPHandlers(router, userUseCase)

	app := &App{
		httpServer: &http.Server{
			Addr:           ":" + port,
			Handler:        router,
			MaxHeaderBytes: 1 << 20,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
		},
	}

	return app
}

func InitDb() *database.Postgres {
	db, err := database.NewClient(viper.GetString("database.uri"))
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func (app *App) Shutdown() error {
	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return app.httpServer.Shutdown(ctx)
}
