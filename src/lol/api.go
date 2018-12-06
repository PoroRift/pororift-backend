package lol

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const API_FILE string = "../api_key"

const URL_SUMMONER = "https://na1.api.riotgames.com/lol/summoner/v4/summoners/by-name/"

var API_KEY string

func Test() string {
	return "Test Package"
}

func init() {
	key, err := readKey()
	if err != nil {
		// fmt.Println(API_KEY, err)
		log.Fatal(err)
	}

	API_KEY = string(key)
	// fmt.Println(API_KEY)
}

func readKey() (string, error) {

	b, err := ioutil.ReadFile(API_FILE)
	if err != nil {
		fmt.Print(err)
		return fmt.Sprintf("Cannot read file: %s", "api_key"), err
	} else {
		// fmt.Println(b) // print byte code (ASCII)
		str := string(b)

		return str, nil
	}
}

func GetSummonerStat(s string) (string, error) {

	url := fmt.Sprintf("%s%s", URL_SUMMONER, s)
	fmt.Println(url)

	// resp, err := http.Get(url)
	// defer resp.Body.Close()
	// if err != nil {
	// 	// handle error
	// 	return "", err
	// }

	// body, err := ioutil.ReadAll(resp.Body)
	// fmt.Println(body)
	// if err != nil {
	// 	return "", err
	// }

	// return string(body), nil

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("X-Riot-Token", API_KEY)
	req.Header.Add("Accept-Charset", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("Accept-Language", "en-US,en;q=0.9")

	fmt.Println(req)
	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil

}

// func makeRequest(method, url string, b *Buffer) {

// 	client := &http.
// }
