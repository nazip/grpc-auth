package main

import (
	"context"
	"flag"
	"log"

	"github.com/nazip/grpc-auth/internal/app"
)

func main() {
	var configFile string

	flag.StringVar(&configFile, "config-file", "local.env", "path to config file")
	flag.Parse()

	ctx := context.Background()

	a, err := app.NewApp(ctx, configFile)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
