package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func AppendHttp(url string) string {
	return fmt.Sprintf("http://%s", url)
}

func ParseUrl(rawUrl string) (width int, height int, url string) {
	splittedUrl := strings.SplitN(rawUrl, "/", 2)

	width, err := strconv.Atoi(splittedUrl[0])
	if err != nil {
		return -1, -1, ""
	}
	height, err = strconv.Atoi(splittedUrl[1])
	if err != nil {
		return -1, -1, ""
	}
	url = splittedUrl[3]

	return
}
