package utils

import (
	"strconv"
	"strings"
)

func ParseUrl(rawUrl string) (width int, height int, url string) {
	splittedUrl := strings.SplitN(rawUrl, "/", 3)

	width, err := strconv.Atoi(splittedUrl[0])
	if err != nil {
		return -1, -1, ""
	}
	height, err = strconv.Atoi(splittedUrl[1])
	if err != nil {
		return -1, -1, ""
	}
	url = splittedUrl[2]

	return
}
