package data

import (
	"testing"
)

// func TestGetPlayer(t *testing.T) {

// 	// t.Log("Given Player does not exists")
// 	// {
// 	// 	GetPlayer("richerthanu")

// 	// }
// }

// t.Log("Given Getting Champion Rotation from Riot")
// 	{
// 		res, err := GetChampRot("na1")
// 		if err != nil {
// 			t.Error("\tShould received champion rotation response", ballotX, err)
// 		}
// 		t.Log("\tShould received champion rotation response", checkMark)

// 		if res.StatusCode == 200 {
// 			t.Log("\tShould received status code 200", checkMark)
// 		} else {
// 			t.Error("\tShould received status code 200", ballotX, res.StatusCode)
// 		}
// 	}

func TestGetPlayer(t *testing.T) {

	Init()
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		// want []byte
	}{
		// TODO: Add test cases.
		{
			name: "Get Player richerthanu",
			args: args{
				name: "richerthanu",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if got := GetPlayer(tt.args.name); !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("GetPlayer() = %v, want %v", got, tt.want)
			// }
			got := GetPlayer(tt.args.name)
			t.Log(string(got))
		})
	}
}

func TestGetMatch(t *testing.T) {
	Init()
	type args struct {
		matchID int
	}
	tests := []struct {
		name string
		args args
		want Match
	}{
		// TODO: Add test cases.
		{
			name: "GetMatch",
			args: args{
				matchID: 2856049818,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if got := GetMatch(tt.args.matchID); !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("GetMatch() = %v, want %v", got, tt.want)
			// }
			got, err := GetMatch(tt.args.matchID)
			t.Log(got, err)
		})
	}
}
