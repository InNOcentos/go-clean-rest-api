package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/InNOcentos/go-clean-rest-api/internal/config"
	"github.com/spf13/viper"

	"github.com/InNOcentos/go-clean-rest-api/internal/app"
)

func main() {
	err := config.Load()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	app := app.NewApp(viper.GetString("server.port"))

	go app.Run()

	quitCh := make(chan os.Signal, 1)

	signal.Notify(quitCh, os.Interrupt, os.Interrupt)

	<-quitCh

	if err = app.Shutdown(); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
