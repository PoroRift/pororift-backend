package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/PoroRift/pororift-backend/lol"
)

type (
	// Match contain match information, and internal components
	match struct {
		GameID int
		mutex  *sync.Mutex
		Data   MatchInfo
	}

	// MatchInfo contain raw data provided by Riot
	MatchInfo struct {
		GameID       int64  `json:"gameId"`
		PlatformID   string `json:"platformId"`
		GameCreation int64  `json:"gameCreation"`
		GameDuration int    `json:"gameDuration"`
		QueueID      int    `json:"queueId"`
		MapID        int    `json:"mapId"`
		SeasonID     int    `json:"seasonId"`
		GameVersion  string `json:"gameVersion"`
		GameMode     string `json:"gameMode"`
		GameType     string `json:"gameType"`
		Teams        []struct {
			TeamID               int           `json:"teamId"`
			Win                  string        `json:"win"`
			FirstBlood           bool          `json:"firstBlood"`
			FirstTower           bool          `json:"firstTower"`
			FirstInhibitor       bool          `json:"firstInhibitor"`
			FirstBaron           bool          `json:"firstBaron"`
			FirstDragon          bool          `json:"firstDragon"`
			FirstRiftHerald      bool          `json:"firstRiftHerald"`
			TowerKills           int           `json:"towerKills"`
			InhibitorKills       int           `json:"inhibitorKills"`
			BaronKills           int           `json:"baronKills"`
			DragonKills          int           `json:"dragonKills"`
			VilemawKills         int           `json:"vilemawKills"`
			RiftHeraldKills      int           `json:"riftHeraldKills"`
			DominionVictoryScore int           `json:"dominionVictoryScore"`
			Bans                 []interface{} `json:"bans"`
		} `json:"teams"`
		Participants []struct {
			ParticipantID             int    `json:"participantId"`
			TeamID                    int    `json:"teamId"`
			ChampionID                int    `json:"championId"`
			Spell1ID                  int    `json:"spell1Id"`
			Spell2ID                  int    `json:"spell2Id"`
			HighestAchievedSeasonTier string `json:"highestAchievedSeasonTier"`
			Stats                     struct {
				ParticipantID                  int  `json:"participantId"`
				Win                            bool `json:"win"`
				Item0                          int  `json:"item0"`
				Item1                          int  `json:"item1"`
				Item2                          int  `json:"item2"`
				Item3                          int  `json:"item3"`
				Item4                          int  `json:"item4"`
				Item5                          int  `json:"item5"`
				Item6                          int  `json:"item6"`
				Kills                          int  `json:"kills"`
				Deaths                         int  `json:"deaths"`
				Assists                        int  `json:"assists"`
				LargestKillingSpree            int  `json:"largestKillingSpree"`
				LargestMultiKill               int  `json:"largestMultiKill"`
				KillingSprees                  int  `json:"killingSprees"`
				LongestTimeSpentLiving         int  `json:"longestTimeSpentLiving"`
				DoubleKills                    int  `json:"doubleKills"`
				TripleKills                    int  `json:"tripleKills"`
				QuadraKills                    int  `json:"quadraKills"`
				PentaKills                     int  `json:"pentaKills"`
				UnrealKills                    int  `json:"unrealKills"`
				TotalDamageDealt               int  `json:"totalDamageDealt"`
				MagicDamageDealt               int  `json:"magicDamageDealt"`
				PhysicalDamageDealt            int  `json:"physicalDamageDealt"`
				TrueDamageDealt                int  `json:"trueDamageDealt"`
				LargestCriticalStrike          int  `json:"largestCriticalStrike"`
				TotalDamageDealtToChampions    int  `json:"totalDamageDealtToChampions"`
				MagicDamageDealtToChampions    int  `json:"magicDamageDealtToChampions"`
				PhysicalDamageDealtToChampions int  `json:"physicalDamageDealtToChampions"`
				TrueDamageDealtToChampions     int  `json:"trueDamageDealtToChampions"`
				TotalHeal                      int  `json:"totalHeal"`
				TotalUnitsHealed               int  `json:"totalUnitsHealed"`
				DamageSelfMitigated            int  `json:"damageSelfMitigated"`
				DamageDealtToObjectives        int  `json:"damageDealtToObjectives"`
				DamageDealtToTurrets           int  `json:"damageDealtToTurrets"`
				VisionScore                    int  `json:"visionScore"`
				TimeCCingOthers                int  `json:"timeCCingOthers"`
				TotalDamageTaken               int  `json:"totalDamageTaken"`
				MagicalDamageTaken             int  `json:"magicalDamageTaken"`
				PhysicalDamageTaken            int  `json:"physicalDamageTaken"`
				TrueDamageTaken                int  `json:"trueDamageTaken"`
				GoldEarned                     int  `json:"goldEarned"`
				GoldSpent                      int  `json:"goldSpent"`
				TurretKills                    int  `json:"turretKills"`
				InhibitorKills                 int  `json:"inhibitorKills"`
				TotalMinionsKilled             int  `json:"totalMinionsKilled"`
				NeutralMinionsKilled           int  `json:"neutralMinionsKilled"`
				TotalTimeCrowdControlDealt     int  `json:"totalTimeCrowdControlDealt"`
				ChampLevel                     int  `json:"champLevel"`
				VisionWardsBoughtInGame        int  `json:"visionWardsBoughtInGame"`
				SightWardsBoughtInGame         int  `json:"sightWardsBoughtInGame"`
				FirstBloodKill                 bool `json:"firstBloodKill"`
				FirstBloodAssist               bool `json:"firstBloodAssist"`
				FirstTowerKill                 bool `json:"firstTowerKill"`
				FirstTowerAssist               bool `json:"firstTowerAssist"`
				FirstInhibitorKill             bool `json:"firstInhibitorKill"`
				FirstInhibitorAssist           bool `json:"firstInhibitorAssist"`
				CombatPlayerScore              int  `json:"combatPlayerScore"`
				ObjectivePlayerScore           int  `json:"objectivePlayerScore"`
				TotalPlayerScore               int  `json:"totalPlayerScore"`
				TotalScoreRank                 int  `json:"totalScoreRank"`
				PlayerScore0                   int  `json:"playerScore0"`
				PlayerScore1                   int  `json:"playerScore1"`
				PlayerScore2                   int  `json:"playerScore2"`
				PlayerScore3                   int  `json:"playerScore3"`
				PlayerScore4                   int  `json:"playerScore4"`
				PlayerScore5                   int  `json:"playerScore5"`
				PlayerScore6                   int  `json:"playerScore6"`
				PlayerScore7                   int  `json:"playerScore7"`
				PlayerScore8                   int  `json:"playerScore8"`
				PlayerScore9                   int  `json:"playerScore9"`
				Perk0                          int  `json:"perk0"`
				Perk0Var1                      int  `json:"perk0Var1"`
				Perk0Var2                      int  `json:"perk0Var2"`
				Perk0Var3                      int  `json:"perk0Var3"`
				Perk1                          int  `json:"perk1"`
				Perk1Var1                      int  `json:"perk1Var1"`
				Perk1Var2                      int  `json:"perk1Var2"`
				Perk1Var3                      int  `json:"perk1Var3"`
				Perk2                          int  `json:"perk2"`
				Perk2Var1                      int  `json:"perk2Var1"`
				Perk2Var2                      int  `json:"perk2Var2"`
				Perk2Var3                      int  `json:"perk2Var3"`
				Perk3                          int  `json:"perk3"`
				Perk3Var1                      int  `json:"perk3Var1"`
				Perk3Var2                      int  `json:"perk3Var2"`
				Perk3Var3                      int  `json:"perk3Var3"`
				Perk4                          int  `json:"perk4"`
				Perk4Var1                      int  `json:"perk4Var1"`
				Perk4Var2                      int  `json:"perk4Var2"`
				Perk4Var3                      int  `json:"perk4Var3"`
				Perk5                          int  `json:"perk5"`
				Perk5Var1                      int  `json:"perk5Var1"`
				Perk5Var2                      int  `json:"perk5Var2"`
				Perk5Var3                      int  `json:"perk5Var3"`
				PerkPrimaryStyle               int  `json:"perkPrimaryStyle"`
				PerkSubStyle                   int  `json:"perkSubStyle"`
			} `json:"stats"`
			Timeline struct {
				ParticipantID      int `json:"participantId"`
				CreepsPerMinDeltas struct {
					Zero10 float64 `json:"0-10"`
				} `json:"creepsPerMinDeltas"`
				XpPerMinDeltas struct {
					Zero10 float64 `json:"0-10"`
				} `json:"xpPerMinDeltas"`
				GoldPerMinDeltas struct {
					Zero10 float64 `json:"0-10"`
				} `json:"goldPerMinDeltas"`
				CsDiffPerMinDeltas struct {
					Zero10 float64 `json:"0-10"`
				} `json:"csDiffPerMinDeltas"`
				XpDiffPerMinDeltas struct {
					Zero10 float64 `json:"0-10"`
				} `json:"xpDiffPerMinDeltas"`
				DamageTakenPerMinDeltas struct {
					Zero10 float64 `json:"0-10"`
				} `json:"damageTakenPerMinDeltas"`
				DamageTakenDiffPerMinDeltas struct {
					Zero10 float64 `json:"0-10"`
				} `json:"damageTakenDiffPerMinDeltas"`
				Role string `json:"role"`
				Lane string `json:"lane"`
			} `json:"timeline"`
		} `json:"participants"`
		ParticipantIdentities []struct {
			ParticipantID int `json:"participantId"`
			Player        struct {
				PlatformID        string `json:"platformId"`
				AccountID         string `json:"accountId"`
				SummonerName      string `json:"summonerName"`
				SummonerID        string `json:"summonerId"`
				CurrentPlatformID string `json:"currentPlatformId"`
				CurrentAccountID  string `json:"currentAccountId"`
				MatchHistoryURI   string `json:"matchHistoryUri"`
				ProfileIcon       int    `json:"profileIcon"`
			} `json:"player"`
		} `json:"participantIdentities"`
	}
)

// Init initilize Match Object and Call Riot's API for that match information
// While Match Object is updating, the "mutex" component is locked
// Will unlock the mutex component once completed
func (m *match) Init() error {

	m.mutex.Lock() // Prevent anyone from editing
	defer m.mutex.Unlock()

	res, err := lol.GetMatchAPI(m.GameID, "na1")
	if err != nil {
		return err
	}

	// var matchInfo MatchInfo
	json.NewDecoder(res.Body).Decode(&m.Data)
	fmt.Println(m.Data)

	// fmt.Println(res.Body)
	// print(res)
	return nil
}

func createMatch(matchID int) *match {
	return &match{
		GameID: matchID,
		mutex:  &sync.Mutex{},
	}
}

func print(resp *http.Response) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// return nil, err
		return
	}
	fmt.Println(string(body))
}
