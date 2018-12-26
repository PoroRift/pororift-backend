package data

import (
	"reflect"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestDBGetPlayer(t *testing.T) {
	db := &DataBase{}
	db.Init()
	t.Log("Given searching for summoner that's not in the database")
	{
		summoner, err := db.GetPlayer("richerthanu")
		// t.Log(*summoner)
		if err != nil {
			t.Error("\tShoud search for summoner ", "richerthanu", ballotX, err)
		}
		t.Log("\tShould search for summoner ", "richerthanu", checkMark)

		// fmt.Println(summoner)
		t.Log("\tWhen Searching for summoner that exist in the database")
		{
			anotherSummoner, _ := db.GetPlayer("richerthanu")
			// var s Summoner = *anotherSummoner
			// t.Log(&s)
			// t.Log(*anotherSummoner)
			check := reflect.DeepEqual(anotherSummoner, summoner) // Check if 2 struct are the same
			// t.Log(reflect.DeepEqual(anotherSummoner, summoner))
			if check {
				t.Log("\t\tShould have the same Object", checkMark)
			} else {
				t.Log("\t\tShould have the same Object", ballotX)
			}
		}
	}

}

// func TestDBGetPlayer(t *testing.T) {

// 	db := &DataBase{
// 		Summoners:      map[string]*Summoner{},
// 		mutex_summoner: &sync.Mutex{},
// 	}
// 	t.Log("Given requesting summoner that does not exists")
// 	{
// 		t.Log("\tWhen I adding new Summoner to DataBase")
// 		{
// 			Summoner, err := db.GetPlayer("richerthanu")

// 			if err != nil {
// 				t.Error("\t\tShould have returned Summoner", err, ballotX)
// 			} else {
// 				t.Log("\t\tShould have returned Summoner", checkMark)
// 				{
// 					if Summoner.Stat.AccountId == "oHvddUUaCL8tXyySt8mJZqfpiE_SIentLjiupC__21DZwg" {
// 						t.Log("\t\tShould have an expected AccountID", Summoner.Stat.AccountId, checkMark)
// 					} else {
// 						t.Error("\t\tShould have an expected AccountID", Summoner.Stat.AccountId, ballotX)
// 					}
// 				}
// 			}

// 			// fmt.Println(summoner.mutex)
// 		}
// 	}
// }

// func TestGetSummonerObject(t *testing.T) {
// 	db := &DataBase{
// 		Summoners:      map[string]*Summoner{},
// 		mutex_summoner: &sync.Mutex{},
// 	}

// 	t.Log("Given try to retrive non existing player from db")
// 	{
// 		summonerObject := db.getSummonerObject("Test")

// 		if summonerObject == nil {
// 			t.Error("\tShould have summoner object created", ballotX)
// 		} else {
// 			t.Log("\tShould have summoner object created", checkMark)
// 		}
// 	}
// }

// func TestActionQueue(t *testing.T) {

// 	actionQueue := &actionQueue.ActionQueue{}
// 	actionQueue.Start()

// 	actionQueue.Wait()
// 	actionQueue.Stop()

// 	// aq := &ActionQueue{
// 	// 	que: make(chan Action),
// 	// }
// 	// db := &DataBase{
// 	// 	ActionQueue: aq,
// 	// }

// 	// fmt.Println(db)
// }
