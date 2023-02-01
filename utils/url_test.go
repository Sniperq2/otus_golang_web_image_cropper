package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		c := NewCache(2)

		c.Set("image1", 100)
		c.Set("image2", 200)
		c.Set("image3", 200)
		image, ok := c.Get("image1")
		require.False(t, ok)
		require.Nil(t, image)
	})
}
