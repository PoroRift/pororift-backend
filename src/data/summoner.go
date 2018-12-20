package data

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/PoroRift/pororift-backend/lol"
)

const OUTDATED_TIME = 5 * 60 * 1000 // 5 minutes

type (
	Summoner struct {
		Name       string
		mutex      *sync.Mutex // Lock this Player while updating
		Stat       *Stat
		Matches    *PlayerMatches
		LastUpdate int64 // unix timestamp of the last update
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

	PlayerMatches struct {
		Matches    []*PlayerMatch `json: matches`
		StartIndex int            `json: startindex`
		EndIndex   int            `json: endIndex`
		TotalGames int            `json: totalGames`
	}

	PlayerMatch struct {
		PlatformId string `json: platformId`
		GameId     int    `json: gameId`
		Champion   int    `json: champion`
		Queue      int    `json: queue`
		Season     int    `json: season`
		Timestamp  int    `json: timestamp`
		Role       string `json: role`
		Lane       string `json: lane`
	}
)

func (s *Summoner) IsOutdated() bool {
	currentTime := time.Now().Unix()

	if currentTime-s.LastUpdate > OUTDATED_TIME {
		return true
	}

	return false

}

// Update summoner information
// Call lol package API
// return error if there is one
func (s *Summoner) Update() error {

	s.mutex.Lock() // Preven anyone from editing
	defer s.mutex.Unlock()

	// Get summoner Stat
	res, err := lol.GetSummonerAPI(s.Name, "na1")

	if err != nil {
		return err
	}

	var stat = &Stat{}
	json.NewDecoder(res.Body).Decode(&stat)
	s.Stat = stat

	// Get Match List
	res, err = lol.GetMatchList("na1", s.Stat.AccountId)

	var matches = &PlayerMatches{}
	json.NewDecoder(res.Body).Decode(&matches)

	s.Matches = matches

	// set updated time
	s.LastUpdate = time.Now().Unix()
	return nil
}

func (pm *PlayerMatches) String() string {
	return fmt.Sprintf("%+v\n", *pm)
}

func (m *PlayerMatch) String() string {
	return fmt.Sprintf("%+v\n", *m)
}

// var requestChan = make(chan )
// "ToString" for type Stat
func (s *Stat) String() string {
	// Print the name of the field along with value
	return fmt.Sprintf("%+v\n", *s)
}
