package data

import (
	"sync"
	"testing"
)

// TODO: Have Generic Message declared as string variables

func TestOutDated(t *testing.T) {
	t.Log("Given that a summoner just created")
	{
		summoner := &Summoner{
			mutex: &sync.Mutex{},
			Name:  "Test",
		}
		if summoner.IsOutdated() {
			t.Log("\tShould have true indicating that this summoner need to be updated", checkMark)
		} else {
			t.Error("\tShould have true indicating that this summoner need to be updated", ballotX)
		}

	}
}

func TestSummonerUpdate(t *testing.T) {
	t.Log("Given that updating summoner")
	{
		summoner := &Summoner{
			mutex: &sync.Mutex{},
			Name:  "richerthanu",
		}

		if err := summoner.Update(); err != nil {
			t.Error("\tShould be updated", ballotX, err)
		} else {
			t.Log("\tShould be updated", checkMark)
			{
				if !summoner.IsOutdated() {
					t.Log("\tShould not be outdated", checkMark)
					{
						if summoner.Stat != nil {
							t.Log("\t\tShould not have empty stat", checkMark)
						} else {
							t.Error("\t\tShould not have empty stat", ballotX)
						}

						if summoner.Matches != nil {
							t.Log("\t\tShould not have empty match list", checkMark)
						} else {
							t.Error("\t\tShould not have empty match list", ballotX)
						}
					}
				} else {
					t.Error("\tShould not be outdated", ballotX)
				}
			}
		}
	}

	t.Log("Given Summoner name not found") // Given Summoner doesn't exists/not found
	{
		summoner := &Summoner{
			mutex: &sync.Mutex{},
			Name:  "jfasdklfjasdkfkjdsalfkjadf",
		}

		err := summoner.Update()

		if err == nil {
			t.Error("\tShould return error", ballotX)
		} else {
			t.Log("\tShould return error", checkMark, err)
		}
		// fmt.Println(err)

		// if err := summoner.Update(); err != nil {
		// 	fmt.Println(err)
		// }

	}
}
