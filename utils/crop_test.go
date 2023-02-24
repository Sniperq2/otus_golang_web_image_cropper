package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCropCore(t *testing.T) {
	t.Run("base test core logic", func(t *testing.T) {
		d, err := os.Getwd()
		if err != nil {
			require.NotNil(t, err, "could not load image")
		}
		f, err := os.Open(fmt.Sprintf("%s/tests_cases/Untitled.jpg", filepath.Dir(d)))
		if err != nil {
			require.NotNil(t, err, "could not crop image")
		}
		defer func() {
			if errIn := f.Close(); err != nil {
				err = errIn
			}
		}()

		customImage, err := CropImage(f, 8, 8)
		if err != nil {
			require.NotNil(t, err, "could not crop image")
		}
		require.Equal(t, 8, customImage.Bounds().Dx())
		require.Equal(t, 8, customImage.Bounds().Dy())
	})
}
