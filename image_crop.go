package main

import (
	"fmt"
	"image_croper/controller"
	"image_croper/utils"
	"net/http"
	"os"
	"strconv"

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

	initConfig := &utils.InitConfig{
		CacheSize: cacheSize,
		CachePath: fmt.Sprintf("%s/cache/", cacheDirectory),
	}

	imageCache := utils.NewCache(initConfig.CacheSize)
	imageCache.Clear() //FIXME: clear cache for now

	r.HandleFunc("/fill/{rest:.*}", controller.Cropper(initConfig)).Methods(http.MethodGet)
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Print("Could not start server application")
	}
}
