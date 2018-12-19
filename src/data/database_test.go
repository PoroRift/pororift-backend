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
			Summoner, _ := db.GetPlayer("richerthanu")

			if Summoner.Stat.AccountId == "oHvddUUaCL8tXyySt8mJZqfpiE_SIentLjiupC__21DZwg" {
				t.Log("\t\tShould have an expected AccountID", Summoner.Stat.AccountId, checkMark)
			} else {
				t.Error("\t\tShould have an expected AccountID", Summoner.Stat.AccountId, ballotX)
			}
			// fmt.Println(summoner.mutex)
		}
	}
}
