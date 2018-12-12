package lol

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestReadKey(t *testing.T) {
	t.Log("Given Read Riot API Key")
	{
		key, err := readKey()

		if err != nil {
			t.Fatal("\tShould be able to read API Key.", ballotX, err)
		}
		t.Log("\tShould be able to read API Key: ", key, checkMark)
	}
}

func TestGetSummonerStat(t *testing.T) {

	res, err := getSummonerStat("richerthanu")

	if err != nil {
		t.Errorf("Error, got: %s", err)
	}
	// fmt.Println(res)

	// fmt.Println(res.StatusCode)
	// log.Println(res)
	assert.NotNil(t, res, "The `response` should not be `nil`")
	assert.Equal(t, http.StatusOK, res.StatusCode)
	// assert.Equal(t, )
	// fmt.Println(string(s))
	// assert.NotNil(t, s, "The `response` should not be `nil`")
	// assert.Equal(t, 4, s, "Expecting `4`")

}

func TestSummonerStat(t *testing.T) {
	stat, err := SummonerStat("richerthanu")
	if err != nil {
		t.Errorf("Error, got: %s", err)
	}

	fmt.Println("\t\t", stat)
}
