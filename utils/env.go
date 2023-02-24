package utils

import (
	"log"
	"os"
)

const webImageCropperCache = "WEB_IMAGE_CROPPER_CACHE"

type InitConfig struct {
	CacheHandle Cache
	CachePath   string
}

func SetConfig() string {
	cacheLength, ok := os.LookupEnv(webImageCropperCache)
	if !ok || len(cacheLength) == 0 {
		// sets default value to 10
		if err := os.Setenv(webImageCropperCache, "10"); err != nil {
			log.Println("No image cropper cache set. Set default to 10")
		}
		return "10"
	}

	return cacheLength
}
