package util

import (
	"fmt"
)

func PublicAssetsLink(filename string) string {
	return fmt.Sprintf(
		"%s/assets/%s",
		GetBaseURL(),
		filename,
	)
}

func GetBaseURL() string {
	return GetEnv("BASE_URL", "http://0.0.0.0:8090")
}
