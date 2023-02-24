package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUrlParse(t *testing.T) {
	t.Run("test url parser", func(t *testing.T) {
		width, height, url := ParseURL("200/300/localhost/image.jpg")
		require.Equal(t, 200, width)
		require.Equal(t, 300, height)
		require.Equal(t, "localhost/image.jpg", url)
	})
	t.Run("test wrong url parser", func(t *testing.T) {
		width, height, url := ParseURL("300/localhost/image.jpg")
		require.Equal(t, -1, width)
		require.Equal(t, -1, height)
		require.Equal(t, "", url)
	})
}
