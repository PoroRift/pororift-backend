package lol

import (
	"fmt"
	"testing"
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
}

func TestGetSummonerStat(t *testing.T) {

	s, err := GetSummonerStat("richerthanu")

	if err != nil {
		t.Errorf("Error, got: %s", err)
	}
	fmt.Println(string(s))
	// assert.NotNil(t, s, "The `response` should not be `nil`")
	// assert.Equal(t, 4, s, "Expecting `4`")

}
