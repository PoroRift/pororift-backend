package repository

import (
	"net/http"
	"os"
)

var baseURL = os.Getenv("RIOT_URL")
var apiKey = os.Getenv("RIOT_API")
var client = &http.Client{}
