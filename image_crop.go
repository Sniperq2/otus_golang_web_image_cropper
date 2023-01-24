package main

import (
	"image_croper/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/fill", controller.Cropper).Methods(http.MethodGet)
	if err := http.ListenAndServe(":3000", r); err != nil {
		return
	}
}
