package controller

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"os"
	"strconv"

	"image_croper/utils"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/oliamb/cutter"
)

func Cropper(config *utils.InitConfig) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		customClient := utils.NewClient()

		makeUrl := utils.AppendHttp(mux.Vars(r)["rest"])
		width, height, url := utils.ParseUrl(makeUrl)
		response, err := customClient.Get(url, r.Header)
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
			Width:  width,
			Height: height,
		})

		if err != nil {
			http.Error(w, "Cannot crop image", http.StatusInternalServerError)
		}

		outBytes := new(bytes.Buffer)
		err = jpeg.Encode(outBytes, croppedImg, &jpeg.Options{
			Quality: 100,
		})

		// create uuid for file naming
		newUUID := uuid.NewString()

		out, _ := os.Create(fmt.Sprintf("%s%s.jpg", config.CachePath, newUUID))
		defer func() {
			if errOut := out.Close(); errOut != nil {
				err = errOut
			}
		}()

		err = jpeg.Encode(out, croppedImg, &jpeg.Options{
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
}
