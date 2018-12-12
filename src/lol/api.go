package lol

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type request func(string) (*http.Request, error)

// This need to be read from config
const URL_NA1 = "https://na1.api.riotgames.com"
const URL_GET_SUMMONER = "/lol/summoner/v4/summoners/by-name/"
const URL_GET_CHAMP_ROT = "/lol/platform/v3/champion-rotations"
const URL_GET_MATCH_LIST = "/lol/match/v4/matchlists/by-account/"

var ERROR_UKNOWN_REGION = errors.New("Unknown Region")

var api_key string
var url_regions map[string]string

func init() {
	// Read Key
	key, err := readKey()
	if err != nil {
		log.Println("Cannot read API key", err)
	}

	api_key = key

	url_regions = map[string]string{
		"na1": URL_NA1,
	}

}

func readKey() (string, error) {
	b, err := ioutil.ReadFile(API_FILE)
	if err != nil {
		// log.Println("Cannot read API Key", err)
		return fmt.Sprintf("Cannot read file: %s", API_FILE), err
	} else {
		str := string(b)

		return str, nil
	}
}

// endpoint: /lol/summoner/v4/summoners/by-name/{summonerName}
// https://developer.riotgames.com/api-methods/#summoner-v4
func getSummonerAPI(summoner, region string) (*http.Response, error) {

	endpoint := fmt.Sprintf("%s%s", URL_GET_SUMMONER, summoner)

	return makeRequest(region, endpoint, func(url string) (*http.Request, error) {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}

		return req, nil
	})

}

// endpoint: /lol/platform/v3/champion-rotations
// https://developer.riotgames.com/api-methods/#champion-v3/GET_getChampionInfo
func getChampRot(region string) (*http.Response, error) {
	endpoint := URL_GET_CHAMP_ROT
	return makeRequest(region, endpoint, func(url string) (*http.Request, error) {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}

		return req, nil
	})
}

// endpoint: /lol/match/v3/matchlists/by-account/{accountId}
// https://developer.riotgames.com/api-methods/#match-v3/GET_getMatchlist
// Options: champion, queue, season, endTime, beginTime, endIndex, beginIndex
// Optional query is not implemented yet
func getMatchList(region, accountId string) (*http.Response, error) {
	endpoint := URL_GET_MATCH_LIST + accountId
	return makeRequest(region, endpoint, func(url string) (*http.Request, error) {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}

		return req, nil
	})
}

func makeRequest(region, endpoint string, fn request) (*http.Response, error) {
	if r, exist := url_regions[region]; exist {
		// Region exists
		// make request
		url := fmt.Sprintf("%s%s", r, endpoint)
		req, err := fn(url)
		if err != nil {
			return nil, err
		}

		// set riot-token
		req.Header.Add("X-Riot-Token", api_key)
		req.Header.Add("Accept-Charset", "application/x-www-form-urlencoded; charset=UTF-8")
		req.Header.Add("Accept-Language", "en-US,en;q=0.9")

		// Do the request
		client := &http.Client{}
		res, err := client.Do(req)

		return res, err

	} else {
		return nil, ERROR_UKNOWN_REGION
	}
}
