package main

import (
	"os"
	"time"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	service := os.Getenv("SERVICE_NAME")
	if service == "" {
		service = "parcelhub"
	}

	sugar.Infof("%s worker started and listening for jobs...", service)

	// Simulate background loop
	for {
		time.Sleep(30 * time.Second)
		sugar.Infof("%s worker heartbeat", service)
	}
}
