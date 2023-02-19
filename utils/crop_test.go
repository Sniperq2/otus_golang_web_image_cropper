package utils

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCropCore(t *testing.T) {
	t.Run("base test core logic", func(t *testing.T) {
		d, err := os.Getwd()
		if err != nil {
			require.NotNil(t, err, "could not load image")
		}
		f, err := os.Open(fmt.Sprintf("%stests_cases/Untitled.jpg", d))
		if err != nil {
			require.NotNil(t, err, "could not crop image")
		}
		customImage, err := CropImage(f, 8, 8)
		if err != nil {
			require.NotNil(t, err, "could not crop image")
		}
		require.Equal(t, 8, customImage.Bounds().Dx())
		require.Equal(t, 8, customImage.Bounds().Dy())
	})
}
