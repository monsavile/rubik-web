package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/monsavile/rubik-web/internal/config"
)

type HTTPServerConfig interface {
	Port() int
}

func scrumbleHandler(res http.ResponseWriter, req *http.Request) {
	result := struct {
		Ok   bool
		Name string
	}{Ok: true, Name: "scrumble"}

	res.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(res).Encode(result); err != nil {
		http.Error(res, "Failed to encode result data", http.StatusInternalServerError)
		return
	}
}

func resolveHandler(res http.ResponseWriter, req *http.Request) {
	result := struct {
		Ok   bool
		Name string
	}{Ok: true, Name: "resolve"}

	res.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(res).Encode(result); err != nil {
		http.Error(res, "Failed to encode result data", http.StatusInternalServerError)
		return
	}
}

func main() {
	var err error

	if err = config.Load(); err != nil {
		log.Fatalf("failed to load config: %s", err)
	}

	var httpServerConfig HTTPServerConfig

	if httpServerConfig, err = config.NewHTTPServerConfig(); err != nil {
		log.Fatalf("failed to get http server config: %s", err)
	}

	router := chi.NewRouter()

	router.Post("/scrumble", scrumbleHandler)
	router.Post("/resolve", resolveHandler)

	addr := fmt.Sprintf(":%d", httpServerConfig.Port())

	if err = http.ListenAndServe(addr, router); err != nil {
		log.Fatal(err)
	}
}
