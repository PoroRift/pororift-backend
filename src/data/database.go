// Compile and return data
// Handling data changes (Synchronously)

package data

import (
	"sync"
)

var players = map[string]string{}

type (
	DataBase struct {
		Test           string
		Summoners      map[string]*Summoner
		Matches        map[int]*Match
		mutex_summoner *sync.Mutex
		mutex_matches  *sync.Mutex
	}
)

func (d *DataBase) GetMatch(matchId int) (*Match, err) {
	match := d.getMatchObject(matchId)

	matchInfo, err := match.GetInfo()

}

func (d *DataBase) GetPlayer(name string) (*Summoner, error) {

	summoner := d.getSummonerObject(name)
	if summoner.IsOutdated() {
		// summoner information need to be updated
		err := summoner.Update()

		return summoner, err
	}

	return summoner, nil

}

func (d *DataBase) getMatchObject(id int) *Match {
	d.mutex_matches.Lock()
	defer d.mutex_matches.Unlock()

	if match, exist := d.Matches[id]; exist {
		// If Match does not exists, create them
		newMatch = &Match{
			mutex:  &sync.Mutex{},
			GameId: id,
		}

		d.Matches[id] = newMatch

		return newMatch
	} else {
		return match
	}
}

// Lock summoner list in the database
// Return summoner if exists otherwise create an empty summoner,
// adds to the summer list then return the new empty summoner
func (d *DataBase) getSummonerObject(name string) *Summoner {
	d.mutex_summoner.Lock()
	defer d.mutex_summoner.Unlock()

	if summoner, exist := d.Summoners[name]; !exist {
		// Summoner doesn't exists, create Summoner
		newSummoner := &Summoner{
			mutex: &sync.Mutex{},
			Name:  name,
		}
		d.Summoners[name] = newSummoner
		return newSummoner
	} else {
		return summoner
	}
}
