package utils

import (
	"log"
	"os"
)

type InitConfig struct {
	CacheSize int
	CachePath string
}

func SetConfig() string {
	cacheLength, ok := os.LookupEnv("WEB_IMAGE_CROPPER_CACHE")
	if !ok || len(cacheLength) == 0 {
		// sets default value to 10
		err := os.Setenv("WEB_IMAGE_CROPPER_CACHE", "10")
		if err != nil {
			log.Println("No image cropper cache set. Set default to 10")
		}
		return "10"
	}

	return cacheLength
}
