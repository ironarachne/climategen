package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/ironarachne/climategen"
	"github.com/ironarachne/random"
)

func getClimate(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var newClimate climategen.Climate

	random.SeedFromString(id)

	newClimate = climategen.Generate()

	json.NewEncoder(w).Encode(newClimate)
}

func getClimateRandom(w http.ResponseWriter, r *http.Request) {
	var newClimate climategen.Climate

	rand.Seed(time.Now().UnixNano())

	newClimate = climategen.Generate()

	json.NewEncoder(w).Encode(newClimate)
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", getClimateRandom)
	r.Get("/{id}", getClimate)

	fmt.Println("Climate Generator API is online.")
	log.Fatal(http.ListenAndServe(":7515", r))
}
