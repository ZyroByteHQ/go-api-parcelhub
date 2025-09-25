package main

import (
	"encoding/json"
	"net/http"
	"os"

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

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		status := map[string]string{
			"status":  "ok",
			"service": service,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(status)
	})

	sugar.Infof("Starting %s API on :8080", service)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		sugar.Fatalf("server failed: %v", err)
	}
}
