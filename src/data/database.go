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
		mutex_summoner *sync.Mutex
		mutex_matches  *sync.Mutex
	}
)

func (d *DataBase) GetPlayer(name string) (*Summoner, error) {

	summoner := d.getSummonerObject(name)
	if summoner.IsOutdated() {
		// summoner information need to be updated
		err := summoner.Update()

		return summoner, err
	}

	return summoner, nil

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
