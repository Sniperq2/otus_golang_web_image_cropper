package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLRUCache(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		c := NewCache(2)

		c.Set("image1", 100)
		c.Set("image2", 200)
		c.Set("image3", 200)
		image, ok := c.Get("image1")
		require.False(t, ok)
		require.Nil(t, image)
	})

	t.Run("pure logic 2", func(t *testing.T) {
		c := NewCache(3)

		c.Set("image1", 100)
		c.Set("image2", 200)
		c.Set("image3", 300)
		c.Get("image3")
		c.Set("image2", 500)
		c.Set("image1", 700)
		c.Get("image2")
		c.Get("image2")
		c.Set("image4", 1000)
		val, ok := c.Get("image3")
		require.False(t, ok)
		require.Nil(t, val)
	})
}
