package main

import (
	"fmt"
	"image_croper/controller"
	"image_croper/utils"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	cacheDirectory, err := os.Getwd()
	if err != nil {
		return
	}

	cacheSize, err := strconv.Atoi(utils.SetConfig())
	if err != nil {
		return
	}

	initConfig := &utils.InitConfig{
		CacheSize: cacheSize,
		CachePath: fmt.Sprintf("%s/cache/", cacheDirectory),
	}
	r.HandleFunc("/fill", controller.Cropper(initConfig)).Methods(http.MethodGet)
	if err := http.ListenAndServe(":3000", r); err != nil {
		return
	}
}
