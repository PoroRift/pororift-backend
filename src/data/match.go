package data

import "sync"

type (
	Match struct {
		GameId int
		mutex  *sync.Mutex
	}
)

func (m *Match) GetInfo() (string, error) {

	return "", error
}
