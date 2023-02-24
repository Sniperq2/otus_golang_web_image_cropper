package utils

import (
	"strconv"
	"strings"
)

func ParseURL(rawURL string) (width int, height int, url string) {
	splittedURL := strings.SplitN(rawURL, "/", 3)

	if len(splittedURL) != 3 {
		return -1, -1, ""
	}

	width, err := strconv.Atoi(splittedURL[0])
	if err != nil {
		return -1, -1, ""
	}
	height, err = strconv.Atoi(splittedURL[1])
	if err != nil {
		return -1, -1, ""
	}
	url = splittedURL[2]

	return
}
