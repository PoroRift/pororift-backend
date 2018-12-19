package lol

import (
	"errors"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestReadKey(t *testing.T) {
	t.Log("Given Read Riot API Key")
	{
		key, err := readKey()
		if err != nil {
			t.Fatal("\tShould be able to read API Key.", ballotX, err)
		}
		t.Log("\tShould be able to read API Key: ", key, checkMark)
	}
}

func TestGetSummonerStat(t *testing.T) {
	t.Log("Given Getting Summoner Stat from Riot")
	{
		res, err := GetSummonerAPI("richerthanu", "na1")
		if err != nil {
			t.Error("\tShould received summoner response", ballotX, err)
		}

		t.Log("\tShhould received summoner responses", checkMark)

		if res.StatusCode == 200 {
			t.Log("\tShould received status code 200", checkMark)
		} else {
			t.Error("\tShould received status code 200", ballotX, res.StatusCode)
		}
	}

	t.Log("Given wrong region")
	{
		_, err := GetSummonerAPI("richerthanu", "na")
		if err != nil {
			if err == errors.New("Unknown Region") {
				t.Log("\tShould recevied UNKNOWN REGION error", checkMark)
			} else {
				t.Error("\tShould received UNKNOWN REGION error", ballotX, err)
			}
		} else {
			t.Error("\tShould received UNKNOWN REGION error but nil instead", ballotX, err)
		}
	}
}

func TestGetChampRotation(t *testing.T) {
	t.Log("Given Getting Champion Rotation from Riot")
	{
		res, err := GetChampRot("na1")
		if err != nil {
			t.Error("\tShould received champion rotation response", ballotX, err)
		}
		t.Log("\tShould received champion rotation response", checkMark)

		if res.StatusCode == 200 {
			t.Log("\tShould received status code 200", checkMark)
		} else {
			t.Error("\tShould received status code 200", ballotX, res.StatusCode)
		}
	}
}

func TestGetMatchList(t *testing.T) {
	t.Log("Given Getting Match List by champion account ID")
	{
		message := "\tShould received match list response"
		res, err := GetMatchList("na1", "oHvddUUaCL8tXyySt8mJZqfpiE_SIentLjiupC__21DZwg")
		if err != nil {
			t.Error(message, ballotX, err)
		}
		t.Log(message, checkMark)

		message_code := "\tShould received status code 200"
		if res.StatusCode == 200 {
			t.Log(message_code, checkMark)
		} else {
			t.Error(message_code, ballotX, res.StatusCode)
		}
	}
}
