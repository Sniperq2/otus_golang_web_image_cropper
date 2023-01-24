package main

import (
	"bytes"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/oliamb/cutter"
)

func cropper(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("https://" + r.URL.Query().Get("img"))
	if err != nil {
		http.Error(w, "Could not read image", http.StatusBadRequest)
	}

	defer func() {
		if errIn := response.Body.Close(); err != nil {
			err = errIn
		}
	}()

	if response.StatusCode != 200 {
		http.Error(w, "Could not load image", http.StatusTeapot)
	}
	img, _, err := image.Decode(response.Body)
	if err != nil {
		http.Error(w, "Cannot decode image", http.StatusInternalServerError)
	}

	croppedImg, err := cutter.Crop(img, cutter.Config{
		Width:  250,
		Height: 300,
	})

	if err != nil {
		http.Error(w, "Cannot crop image", http.StatusInternalServerError)
	}

	outBytes := new(bytes.Buffer)
	err = jpeg.Encode(outBytes, croppedImg, &jpeg.Options{
		Quality: 100,
	})

	if err != nil {
		http.Error(w, "Cannot crop image", http.StatusInternalServerError)
	}
	w.Header().Add("Content-Type", "image/jpeg")
	w.Header().Add("Content-Length", strconv.Itoa(len(outBytes.Bytes())))
	if _, err := w.Write(outBytes.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/fill", cropper).Methods(http.MethodGet)
	if err := http.ListenAndServe(":3000", r); err != nil {
		return
	}
}
