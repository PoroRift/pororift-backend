// Compile and return data
// Handling data changes (Synchronously)

package data

import (
	"errors"
	"sync"

	"github.com/PoroRift/pororift-backend/actionQueue"
)

var players = map[string]string{}

type (
	// Instance represent an instance of DataBase that contain compiled data
	Instance struct {
		MatchQueue    *actionQueue.ActionQueue
		SummonerQueue *actionQueue.ActionQueue
		Summoners     map[string]*summoner
		Matches       map[int]*match
		mutexSummoner *sync.Mutex
		mutexMatches  *sync.Mutex
	}

	// Request represent a request(action) to create/update data.
	// Request will be send to the action que of that instance.
	Request struct {
		// Respond is an interface object that can hold whatever is defined
		// as the respond from the from the request
		Respond actionQueue.Object
		// A WaitGroup to check if the request is processed and finished
		wg  sync.WaitGroup
		err error
		// action defined what action the request will perform
		action func() (actionQueue.Object, error)
	}
)

// Run defined how the request will be executed. (Not action
// the request wil be performing)
// Method Set for actionQueue.Request interface
// This function will be run by a different thread (Goroutine, ActionQueue.consume())
func (r *Request) Run() {
	defer r.wg.Done() // finished processing

	Object, err := r.action()

	r.Respond = Object // Stored the result of the action
	r.err = err        // Stored the error of the action
}

// Init function initialize the Instance and all it components.
// Call it repeatedly will result in a new component
// TODO: Check if the instance has been called, prevent from reset
// TODO: Create Reset Function
func (db *Instance) Init() {
	db.SummonerQueue = &actionQueue.ActionQueue{}
	db.SummonerQueue.Start()
	db.MatchQueue = &actionQueue.ActionQueue{}
	db.MatchQueue.Start()

	db.Summoners = map[string]*summoner{}
	db.Matches = map[int]*match{}
}

// GetPlayer search the information of that player.
// Store every name that is search including the unknown or non-existing player
// to prevent spam.
// If the player doesn't exist or need to be update, create a request struct
// and put the request to update/create summoner into action queue
func (db *Instance) getPlayer(name string) (*summoner, error) {

	if summoner, exist := db.Summoners[name]; exist {
		// TODO: Check if summoner is outdated and need to be updated
		if summoner.IsOutdated() {
			// summoenr information need to be updated
			request := &Request{
				action: func() (actionQueue.Object, error) {
					err := summoner.Update()
					return nil, err
				},
			}
			request.wg.Add(1)             // Indicate the request is processing
			db.SummonerQueue.Add(request) // Put the request into queue

			request.wg.Wait() // (This thread) wait for request to be finished

			return summoner, request.err
		}
		return summoner, nil

	}

	// Create a new summoner and get information
	request := &Request{
		action: func() (actionQueue.Object, error) {
			// If previous Request(Action) already update summoner
			if summoner, exist := db.Summoners[name]; exist {
				return summoner, nil
			}

			newSummoner := createSummoner(name)

			err := newSummoner.Update()
			db.Summoners[name] = newSummoner
			// fmt.Println("newSummoner: " + reflect.TypeOf(newSummoner).String())

			return newSummoner, err
		},
	}

	request.wg.Add(1)             // Indicate the request is processing
	db.SummonerQueue.Add(request) // Put the request into queue

	request.wg.Wait() // (This thread) wait for request to be finished

	// Converting resopnd(interface) into pointer summoner
	sum, ok := request.Respond.(*summoner)

	if !ok {
		return nil, errors.New("Interface conversion error")
	}
	// fmt.Println(sum, request.err)
	return sum, request.err

	// fmt.Println(sum, ok)
	// var s *Summoner = reflect.ValueOf(request.Respond).Interface().(*Summoner)
	// s := reflect.ValueOf(request.Respond).Interface().(*Summoner)

	// s.Name = "Not Rich"
	// return s, request.err

}

// GetMatch return match information of the provided match ID
func (db *Instance) getMatch(matchID int) (*match, error) {
	if match, exist := db.Matches[matchID]; exist {
		// If match exist, return match information
		return match, nil
	}

	/** Match doesn't exists **/
	request := &Request{
		action: func() (actionQueue.Object, error) {
			if match, exist := db.Matches[matchID]; exist {
				// If previous request already compiled match information
				return match, nil
			}

			newMatch := createMatch(matchID)
			db.Matches[matchID] = newMatch
			err := newMatch.Init()
			return newMatch, err
		},
	}
	request.wg.Add(1)
	db.MatchQueue.Add(request)
	request.wg.Wait()

	m, ok := request.Respond.(*match)

	if !ok {
		return nil, errors.New("Interface convertion error")
	}
	// fmt.Println(m.Data)
	return m, request.err
}
