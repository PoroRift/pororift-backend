package lol

import (
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadKey(t *testing.T) {
	key, err := readKey()
	// fmt.Println(key)
	if err != nil {
		t.Errorf("Error, got: %s", err)
	}

	if key == "" {
		t.Errorf("Error, got :%s", key)
	}

	fmt.Println()
}

func TestGetSummonerStat(t *testing.T) {

	res, err := getSummonerStat("richerthanu")

	if err != nil {
		t.Errorf("Error, got: %s", err)
	}
	// fmt.Println(res)

	// fmt.Println(res.StatusCode)
	log.Println(res)
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

	fmt.Println(stat)
}
