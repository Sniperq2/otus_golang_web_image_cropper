package controller

import (
	"fmt"
	"image_croper/utils"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestCropper(t *testing.T) {
	t.Run("integration test", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/fill/8/8/localhost/Untitled.jpg", nil)
		w := httptest.NewRecorder()
		tempCache := utils.NewCache(10)
		cacheDirectory, err := os.Getwd()
		if err != nil {
			t.Error("could not get program executable path")
		}

		Cropper(&utils.InitConfig{
			CacheHandle: tempCache,
			CachePath:   fmt.Sprintf("%s/cache/", filepath.Dir(cacheDirectory)),
		})(w, req)

		res := w.Result()

		defer func() {
			if errIn := res.Body.Close(); errIn != nil {
				err = errIn
			}
		}()

		if res.StatusCode != 200 {
			t.Error("could not get result")
		}
	})
}
