package utils

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestENVCache(t *testing.T) {
	t.Run("base test env config", func(t *testing.T) {
		envConfig := SetConfig()
		require.Equal(t, "10", envConfig)
	})

	t.Run("test env config", func(t *testing.T) {
		if err := os.Setenv("WEB_IMAGE_CROPPER_CACHE", "20"); err != nil {
			log.Println("No image cropper cache set. Set default to 10")
		}
		envConfig := SetConfig()
		require.Equal(t, "20", envConfig)
	})
}
