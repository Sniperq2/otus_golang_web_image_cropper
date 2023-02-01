package utils

import "fmt"

func AppendHttp(url string) string {
	return fmt.Sprintf("http://%s", url)
}
