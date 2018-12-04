// Package data provide and compile all data that the handler(from client) requested
package data

import (
	"encoding/json"
	"fmt"
	"log"
)

var db Instance

func Init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// db.Test = "Test String"
	db = Instance{}
	db.Init()
}

// Look up player by name
// Return player information
func GetPlayer(name string) []byte {
	fmt.Println(name)
	summoner, err := db.getPlayer(name)
	b, err := json.Marshal(summoner)
	fmt.Println(b, err)
	return b

	// fmt.Println(db.Test)
}

// GetMatch .
func GetMatch(matchID int) (*Match, error) {
	matchInfo, err := db.getMatch(matchID)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// fmt.Println(matchInfo.Data.Participants)
	participants := matchInfo.Data.Participants

	match := Match{}
	match.Summoners = make([]Summoner, len(participants))
	for i := 0; i < len(participants); i++ {
		// participant := (*participants)[i]

		player := participants[i]
		identities := (matchInfo.Data.ParticipantIdentities)[i]
		platform := identities.Player.PlatformID
		accountID := identities.Player.AccountID
		name := identities.Player.SummonerName
		summonerID := identities.Player.SummonerID
		spell1 := player.Spell1ID
		spell2 := player.Spell2ID
		championID := player.ChampionID
		s := &Summoner{
			SummonerName: name,
			AccountID:    accountID,
			TeamID:       player.TeamID,
			Win:          player.Stats.Win,
			Rank:         player.HighestAchievedSeasonTier,
			KDA: KDA{
				Kill:   player.Stats.Kills,
				Death:  player.Stats.Deaths,
				Assist: player.Stats.Assists,
			},
			Champion: ChampIcon{
				ChampionIcon: Icon{
					URL:     "",
					AltText: string(championID),
				}, // TODO
				ChampionName: string(championID), // TODO
			},
			FirstSpellIcon: Icon{
				URL:     "",
				AltText: string(spell1),
			},
			SecondSpellIcon: Icon{
				URL:     "",
				AltText: string(spell2),
			},
		}
		log.Println(platform, accountID, name, summonerID)
		// parID := participants[i].ParticipantID
		// *participants.
		// match.Summoners[i] = Summoner{
		// 	SummonerName: "",
		// }
		match.Summoners[i] = *s
	}

	return &match, nil
}

type (
	Match struct {
		Summoners []Summoner `json:"summonerDetails"`
	}

	Summoner struct {
		SummonerName    string    `json:"summonerName"`
		KDA             KDA       `json:"summonerKDA"`
		Rank            string    `json:"summonerRank"`
		AccountID       string    `json:"accountId"`
		TeamID          int       `json:"TeamID"`
		Win             bool      `json:"win"`
		Champion        ChampIcon `json:"champion"`
		FirstSpellIcon  Icon      `json:"firstSpellIcon"`
		SecondSpellIcon Icon      `json:"secondSpellIcon"`
	}

	ChampIcon struct {
		ChampionName string `json:"championName"`
		ChampionIcon Icon   `json:"championIcon"`
	}

	KDA struct {
		Kill   int `json:"kill"`
		Death  int `json:"death"`
		Assist int `json:"assist"`
	}

	Icon struct {
		URL     string `json:"url"`
		AltText string `json:"alt"`
	}
)
