// Compile and return data
// Handling data changes (Synchronously)

package data

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/PoroRift/pororift-backend/actionQueue"
)

var players = map[string]string{}

type (
	DataBase struct {
		ActionQueue    *actionQueue.ActionQueue
		Summoners      map[string]*Summoner
		Matches        map[int]*Match
		mutex_summoner *sync.Mutex
		mutex_matches  *sync.Mutex
	}

	Request struct {
		Respond *actionQueue.Object
		err     error
		action  func() (actionQueue.Object, error)
	}
)

// Method Set for actionQueue.Request interface
func (r *Request) Run() {
	Object, err := r.action()
	// fmt.Println(reflect.TypeOf(Object).String())
	// fmt.Println(&Object)
	r.Respond = &Object
	// fmt.Println(reflect.TypeOf(r.Respond).String())
	// fmt.Println(r.Respond)
	r.err = err
	// fmt.Println(*r.Respond)
	// fmt.Println(reflect.TypeOf(*r.Respond).String())

}

func (db *DataBase) Init() {
	db.ActionQueue = &actionQueue.ActionQueue{}
	db.ActionQueue.Start()

	db.Summoners = map[string]*Summoner{}
	db.Matches = map[int]*Match{}
}

// type Update func()()

// func (a *ActionGetPlayer) Action(fn Update) {

// }

// func (a *ActionGetPlayer) Action

// func (d *DataBase) GetMatch(matchId int) (*Match, err) {
// 	match := d.getMatchObject(matchId)

// 	matchInfo, err := match.GetInfo()

// }

func (d *DataBase) GetPlayer(name string) (*Summoner, error) {

	// summoner := d.getSummonerObject(name)
	// if summoner.IsOutdated() {
	// 	// summoner information need to be updated
	// 	err := summoner.Update()

	// 	return summoner, err
	// }

	// return summoner, nil

	if summoner, exist := d.Summoners[name]; exist {
		return summoner, nil
	} else {
		// Either response with nil first or wait

		// Put in action Queue
		// action := func() (*Summoner, error) {

		// 	// if other action in the queue already created summoner
		// 	if summoner, exist := d.Summoners[name]; exist {
		// 		return summoner, nil
		// 	}
		// 	d.Summoners[name] = &Summoner{
		// 		Name:  name,
		// 		mutex: &sync.Mutex{},
		// 	}

		// 	return d.Summoners[name], nil
		// }

		action := func() (actionQueue.Object, error) {
			if summoner, exist := d.Summoners[name]; exist { // If previous Action already update summoner
				return summoner, nil
			}

			newSummoner := createSummoner(name)

			err := newSummoner.Update()
			d.Summoners[name] = newSummoner
			return newSummoner, err
		}

		request := &Request{
			action: action,
		}
		d.ActionQueue.Add(request)
		d.ActionQueue.Wait() // wrong, need to change to request lock
		// fmt.Println(reflect.TypeOf(request.Respond).String())
		// t := request.Respond.(Summoner)
		// t := &request.Respond
		// fmt.Println(t.Name)
		// str, ok := request.Respond.(Summoner)
		// fmt.Println(str, ok)
		// type2, panic := request.Respond() //Summoner.(request.Respond)
		// fmt.Println(type2, panic)
		// return Summoner(request.Respond), request.err
		// return request.Respond, request.err
		fmt.Println(reflect.TypeOf((*request.Respond)).String())
		fmt.Println(*request.Respond)
		// var s Summoner = *request.Respond
		// var s Summoner = *request.Respond.(Summoner)
		// var s *Summoner = *request.Respond
		// fmt.Println(reflect.TypeOf(request.Respond).String())
		// fmt.Println(reflect.TypeOf(&request.Respond).String())
		// var s *Summoner = &request.Respond
		return &Summoner{}, nil
	}

}

// func (d *DataBase) getMatchObject(id int) *Match {
// 	d.mutex_matches.Lock()
// 	defer d.mutex_matches.Unlock()

// 	if match, exist := d.Matches[id]; !exist {
// 		// If Match does not exists, create them
// 		newMatch = &Match{
// 			mutex:  &sync.Mutex{},
// 			GameId: id,
// 		}

// 		d.Matches[id] = newMatch

// 		return newMatch
// 	} else {
// 		return match
// 	}
// }

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
