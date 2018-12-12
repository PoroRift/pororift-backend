package lol

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type request struct {
	req map[string]string
}

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

// TODo: Error Checking
func getSummonerStat(s string) (*http.Response, error) {

	url := fmt.Sprintf("%s%s", URL_SUMMONER, s)
	// fmt.Println(url)

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

	// Prepare request
	req, err := http.NewRequest("GET", url, nil)

	// Set headers
	req.Header.Add("X-Riot-Token", API_KEY)
	req.Header.Add("Accept-Charset", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("Accept-Language", "en-US,en;q=0.9")

	// Do the request
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	// fmt.Println(reflect.TypeOf(resp))
	// return resp, nil
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return nil, err
	// }

	// fmt.Println(string(body))

	// result := new(Stat)
	// // var result map[string]interface{}
	// json.NewDecoder(resp.Body).Decode(&result)
	// // result := resp.Body
	// log.Println(result)

	// fmt.Println(reflect.TypeOf(result))
	return resp, nil

	// fmt.Println(string(body))
	// return string(body), nil

}

func SummonerStat(name string) (*Stat, error) {
	res, error := getSummonerStat(name)

	if error != nil {
		return nil, error
	}

	var stat = &Stat{}
	json.NewDecoder(res.Body).Decode(&stat)
	// log.Println(stat)

	return stat, nil
}

// func makeRequest(method, url string, b *Buffer) {

// 	client := &http.Client{}

// 	req, err := http.NewRequest("method", url, nil)
// }
