package data

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/PoroRift/pororift-backend/lol"
)

const outDatedTime = 5 * 60 * 1000 // 5 minutes

type (

	// Summoner contains player's stat and information
	summoner struct {
		Name string
		// Lock this Player while updating
		mutex   *sync.Mutex
		Stat    *Stat
		Matches *summonerMatch
		// unix timestamp of the last update
		LastUpdate int64
	}

	summonerMatch struct {
		Matches    []*matchInfo `json:"matches"`
		StartIndex int          `json:"startindex"`
		EndIndex   int          `json:"endIndex"`
		TotalGames int          `json:"totalGames"`
	}
	matchInfo []struct {
		PlatformID string `json:"platformId"`
		GameID     int    `json:"gameId"`
		Champion   int    `json:"champion"`
		Queue      int    `json:"queue"`
		Season     int    `json:"season"`
		Timestamp  int    `json:"timestamp"`
		Role       string `json:"role"`
		Lane       string `json:"lane"`
	}
	// Stat is what Riot's API return regarding player information
	Stat struct {
		ID           string `json:"id"`
		AccountID    string `json:"accountId"`
		Puuid        string `json:"puuid"`
		Name         string `json:"name"`
		ProfileIcon  int    `json:"profileIconId"`
		RevisionDate int    `json:"revisionDate"`
		Lvl          int    `json:"summonerLevel"`
	}
	// PlayerMatches List of the summoner's recent matches
)

// IsOutdated check the last timestamp.
// Return true if timestamp exceeded the set amount by
// summoner's LastUpdate.
func (s *summoner) IsOutdated() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	currentTime := time.Now().Unix()

	if currentTime-s.LastUpdate > outDatedTime {
		return true
	}

	return false

}

// Update summoner information
// Call lol package API
// return error if there is one
func (s *summoner) Update() error {

	s.mutex.Lock() // Prevent anyone from editing
	defer s.mutex.Unlock()

	// Get summoner Stat
	res, err := lol.GetSummonerAPI(s.Name, "na1")
	if err != nil {
		return err
	}

	json.NewDecoder(res.Body).Decode(&s.Stat)

	// Get Match List
	res, err = lol.GetMatchList("na1", s.Stat.AccountID)

	json.NewDecoder(res.Body).Decode(&s.Matches)

	// set updated time
	s.LastUpdate = time.Now().Unix()
	return nil
}

// func (pm *PlayerMatches) String() string {
// 	return fmt.Sprintf("%+v\n", *pm)
// }

// func (m *PlayerMatch) String() string {
// 	return fmt.Sprintf("%+v\n", *m)
// }

// // var requestChan = make(chan )
// // "ToString" for type Stat
// func (s *Stat) String() string {
// 	// Print the name of the field along with value
// 	return fmt.Sprintf("%+v\n", *s)
// }

func createSummoner(name string) *summoner {
	return &summoner{
		Name:  name,
		mutex: &sync.Mutex{},
	}
}
