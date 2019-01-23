package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var baseURL, apiKey string
var client = &http.Client{}

//Initialize initializes objects given a repository
func Initialize() {
	apiKey = os.Getenv("RIOT_API")
	baseURL = os.Getenv("RIOT_URL")
}

func getBaseURL(extension string) string {
	return fmt.Sprintf("%s%s?api_key=%s", baseURL, extension, apiKey)
}

func callHandler(req *http.Request, err error) (*interface{}, error) {
	var obj interface{}

	res, fetchErr := client.Do(req)

	if fetchErr != nil {
		return nil, fetchErr
	}

	if res.StatusCode != 200 {
		return nil, errors.New(res.Status)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {

		return nil, readErr
	}

	jsonErr := json.Unmarshal(body, &obj)

	if jsonErr != nil {
		return nil, jsonErr
	}

	return &obj, nil
}
