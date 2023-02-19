package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"image_croper/controller"
	"image_croper/utils"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	r := mux.NewRouter()
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	cacheDirectory, err := os.Getwd()
	if err != nil {
		log.Print("could not get program executable path")
	}

	cacheSize, err := strconv.Atoi(utils.SetConfig())
	if err != nil {
		log.Print("Could not set proper config")
	}
	cacheHandle := utils.NewCache(cacheSize)

	initConfig := &utils.InitConfig{
		CacheHandle: cacheHandle,
		CachePath:   fmt.Sprintf("%s/cache/", cacheDirectory),
	}

	r.HandleFunc("/fill/{rest:.*}", controller.Cropper(initConfig)).Methods(http.MethodGet)

	server := &http.Server{
		Addr:              ":3000",
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           r,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Print("Could not start server application")
	}
}
