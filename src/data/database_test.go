package data

import (
	"sync"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestDBGetPlayer(t *testing.T) {

	db := &DataBase{
		Summoners:      map[string]*Summoner{},
		mutex_summoner: &sync.Mutex{},
	}
	t.Log("Given requesting summoner that does not exists")
	{
		t.Log("\tWhen I adding new Summoner to DataBase")
		{
			Summoner, err := db.GetPlayer("richerthanu")

			if err != nil {
				t.Error("\t\tShould have returned Summoner", err, ballotX)
			} else {
				t.Log("\t\tShould have returned Summoner", checkMark)
				{
					if Summoner.Stat.AccountId == "oHvddUUaCL8tXyySt8mJZqfpiE_SIentLjiupC__21DZwg" {
						t.Log("\t\tShould have an expected AccountID", Summoner.Stat.AccountId, checkMark)
					} else {
						t.Error("\t\tShould have an expected AccountID", Summoner.Stat.AccountId, ballotX)
					}
				}
			}

			// fmt.Println(summoner.mutex)
		}
	}
}

func TestGetSummonerObject(t *testing.T) {
	db := &DataBase{
		Summoners:      map[string]*Summoner{},
		mutex_summoner: &sync.Mutex{},
	}

	t.Log("Given try to retrive non existing player from db")
	{
		summonerObject := db.getSummonerObject("Test")

		if summonerObject == nil {
			t.Error("\tShould have summoner object created", ballotX)
		} else {
			t.Log("\tShould have summoner object created", checkMark)
		}
	}
}
