// Compile and return data
// Handling data changes (Synchronously)

package data

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/PoroRift/pororift-backend/lol"
)

var players = map[string]string{}

type (
	DataBase struct {
		Test           string
		Summoners      map[string]*Summoner
		mutex_summoner *sync.Mutex
		mutex_matches  *sync.Mutex
	}

	Summoner struct {
		mutex   *sync.Mutex // lock this Player while updating
		Stat    *Stat
		Matches *Matches
	}

	Stat struct {
		Id           string `json:"id"`
		AccountId    string `json:"accountId"`
		Puuid        string `json:"puuid"`
		Name         string `json:"name"`
		ProfileIcon  int    `json:"profileIconId"`
		RevisionDate int    `json:"revisionDate"`
		Lvl          int    `json:"summonerLevel"`
	}

	// TODO: Break into another file
	Matches struct {
		Matches    []Match `json: matches`
		StartIndex int     `json: startindex`
		EndIndex   int     `json: endIndex`
		TotalGames int     `json: totalGames`
	}

	Match struct {
		PlatformId string `json: platformId`
		GameId     int    `json: gameId`
		Champion   int    `json: champion`
		Queue      int    `json: queue`
		Season     int    `json:season`
		Timestamp  int    `json:timestamp`
		Role       string `json: role`
		Lane       string `json:lane`
	}
)

func (d *DataBase) GetPlayer(name string) (*Summoner, error) {
	// Check if the player exists in player list

	// TODO: Maybe break into another function
	d.mutex_summoner.Lock() // Prevent Other from Updating

	if summoner, exist := d.Summoners[name]; exist {
		defer d.mutex_summoner.Unlock()
		// Player exists
		return summoner, nil
	} else {
		// Player does not exists and need to be created

		newSummoner := &Summoner{
			mutex: &sync.Mutex{},
		}

		d.Summoners[name] = newSummoner
		d.mutex_summoner.Unlock()

		newSummoner.mutex.Lock()
		defer newSummoner.mutex.Unlock()

		// Get Summoner Stat TODO: Break into a function
		res, err := lol.GetSummonerAPI(name, "na1")

		if err != nil {
			return nil, err
		}

		var stat = &Stat{}

		json.NewDecoder(res.Body).Decode(&stat)
		newSummoner.Stat = stat

		// Get Match List TODO: Break into a function
		res, err = lol.GetMatchList("na1", newSummoner.Stat.AccountId)

		var matches = &Matches{}
		json.NewDecoder(res.Body).Decode(&matches)

		newSummoner.Matches = matches
		fmt.Println(newSummoner)
		return d.Summoners[name], nil

		// return nil, &Summoner{}
	}
}

// var requestChan = make(chan )
// "ToString" for type Stat
func (s *Stat) String() string {
	// Print the name of the field along with value
	return fmt.Sprintf("%+v\n", *s)
}

func (s *Matches) String() string {
	return fmt.Sprintf("%+v\n", *s)
}
