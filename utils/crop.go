package utils

import (
	"errors"
	"image"
	"io"

	"github.com/oliamb/cutter"
)

func CropImage(webImage io.ReadCloser, width int, height int) (image.Image, error) {
	img, _, err := image.Decode(webImage)
	if err != nil {
		return nil, errors.New("cannot decode image")
	}

	croppedImg, err := cutter.Crop(img, cutter.Config{
		Width:  width,
		Height: height,
	})

	if err != nil {
		return nil, errors.New("cannot crop image")
	}

	return croppedImg, nil
}
