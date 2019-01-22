package repository

import (
	"fmt"
	"net/http"
	"os"
)

var baseURL, apiKey string

var client = &http.Client{}

func Initialize() {
	apiKey = os.Getenv("RIOT_API")
	baseURL = os.Getenv("RIOT_URL")
}

func getBaseUrl(extension string) string {
	return fmt.Sprintf("%s%s?api_key=%s", baseURL, extension, apiKey)
}
