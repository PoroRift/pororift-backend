package data

import (
	"reflect"
	"sync"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestDBGetPlayer(t *testing.T) {
	db := &Instance{}
	db.Init()
	t.Log("Given searching for summoner that's not in the database")
	{
		summoner, err := db.getPlayer("richerthanu")
		// t.Log(*summoner)
		if err != nil {
			t.Error("\tShoud search for summoner ", "richerthanu", ballotX, err)
		}
		t.Log("\tShould search for summoner ", "richerthanu", checkMark)

		// fmt.Println(summoner)
		t.Log("\tWhen Searching for summoner that exist in the database")
		{
			anotherSummoner, _ := db.getPlayer("richerthanu")
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

func TestGetTwoPlayer(t *testing.T) {
	db := &Instance{}
	db.Init()
	t.Log("Given requesting 2 player at the same time")
	{
		var wg sync.WaitGroup
		wg.Add(2)
		var s1, s2 *summoner
		go func() {
			sum, _ := db.getPlayer("richerthanu")
			s1 = sum
			wg.Done()
		}()

		go func() {
			sum, _ := db.getPlayer("rubberice")
			s2 = sum
			wg.Done()
		}()

		wg.Wait()

		// t.Log(*s1, *s2)
		t.Log(s1)
		if s1 != nil && s2 != nil {
			t.Log("\tShould recevied 2 summoners objects", checkMark)
		} else {
			t.Error("\tShould recevied 2 summoners objects", ballotX)
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

func TestInstance_GetPlayer(t *testing.T) {

	instance := &Instance{}
	instance.Init()

	type args struct {
		name string
	}
	tests := []struct {
		name     string
		Instance *Instance
		args     args
		want     string // Acount ID
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name:     "Case 1: Get Summoner `richerthanu`",
			Instance: instance,
			args: args{
				name: "richerthanu",
			},
			want:    "oHvddUUaCL8tXyySt8mJZqfpiE_SIentLjiupC__21DZwg",
			wantErr: false,
		},
		{
			name:     "Case 2: Get Summoner rubberice",
			Instance: instance,
			args: args{
				name: "rubberice",
			},
			want:    "eeZuGNed2rhONqLW7wV75oqPSkCoZdnLgodHcak0fhFBgg",
			wantErr: false,
		},
		{
			name:     "Case 3: Get Summoner `richerthanu`",
			Instance: instance,
			args: args{
				name: "richerthanu",
			},
			want:    "oHvddUUaCL8tXyySt8mJZqfpiE_SIentLjiupC__21DZwg",
			wantErr: false,
		},
		{
			name:     "Case 4: Get Summoner rubberice",
			Instance: instance,
			args: args{
				name: "rubberice",
			},
			want:    "eeZuGNed2rhONqLW7wV75oqPSkCoZdnLgodHcak0fhFBgg",
			wantErr: false,
		},
		{
			name:     "case 5: Unknown Summoner",
			Instance: instance,
			args: args{
				name: "adklfjasdkldfas",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := tt.Instance
			got, err := db.getPlayer(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Instance.GetPlayer() error = %v, wantErr %v %v", err, tt.wantErr, ballotX)
				return
			}
			// t.Log(got, err)
			if err == nil && !reflect.DeepEqual(got.Stat.AccountID, tt.want) {
				t.Errorf("Instance.GetPlayer() = %v, want %v", got, tt.want)
			}
			t.Logf("Instance.GetPlayer() %v success %v", tt.args.name, checkMark)
		})
	}

	if len(instance.Summoners) == 3 {
		t.Logf("Should have 3 Summoner Objects in list %v", checkMark)
	} else {
		t.Errorf("Should have 3 Summoner Objects in list instead has %v %v", len(instance.Summoners), ballotX)
	}
}

func TestInstance_getMatch(t *testing.T) {
	// type fields struct {
	// 	MatchQueue    *actionQueue.ActionQueue
	// 	SummonerQueue *actionQueue.ActionQueue
	// 	Summoners     map[string]*summoner
	// 	Matches       map[int]*match
	// 	mutexSummoner *sync.Mutex
	// 	mutexMatches  *sync.Mutex
	// }
	instance := &Instance{}
	instance.Init()

	type args struct {
		matchID int
	}
	tests := []struct {
		name string
		// fields  fields
		args args
		// want    *match
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Get Match Information",
			args: args{
				matchID: 2856049818,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := instance
			got, err := db.getMatch(tt.args.matchID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Instance.getMatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("Instance.getMatch() = %v, want %v", got, tt.want)
			// }
			t.Log(got)
		})
	}
}
