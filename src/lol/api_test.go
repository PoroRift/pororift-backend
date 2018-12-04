package lol

import (
	"net/http"
	"reflect"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func Test_readKey(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Read API Key",
			want:    "Test",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := readKey()
			if (err != nil) != tt.wantErr {
				t.Errorf("readKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if got != tt.want {
			// 	t.Errorf("readKey() = %v, want %v", got, tt.want)
			// }
		})
	}
}

func TestGetSummonerAPI(t *testing.T) {
	type args struct {
		summoner string
		region   string
	}
	tests := []struct {
		name    string
		args    *args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Call Summoner api, checking status code",
			args: &args{
				summoner: "richerthanu",
				region:   "na1",
			},
			want:    200,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetSummonerAPI(tt.args.summoner, tt.args.region)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSummonerAPI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("GetSummonerAPI() = %v, want %v", got, tt.want)
			// }
			if got.StatusCode == tt.want {
				t.Logf("GetSummonerAPI() return status code %v %v", got.StatusCode, checkMark)
			} else {
				t.Errorf("GetSummonerAPI() return incorrect status code %v %v", got.StatusCode, ballotX)
				return
			}
		})
	}
}

func TestGetChampRot(t *testing.T) {
	type args struct {
		region string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Get Champion rotation of NA1, checking status code",
			args: args{
				region: "na1",
			},
			want:    200,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetChampRot(tt.args.region)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetChampRot() error = %v, wantErr %v %v", err, tt.wantErr, ballotX)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("GetChampRot() = %v, want %v", got, tt.want)
			// }

			if got.StatusCode != tt.want {
				t.Errorf("GetChampRot() = %v, want %v %v", got, tt.want, ballotX)
			}
			t.Logf("Checking Status Code 200 %v", checkMark)
		})
	}
}

func TestGetMatchList(t *testing.T) {
	type args struct {
		region    string
		accountID string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Get List of Match, checking status code",
			args: args{
				region:    "na1",
				accountID: "oHvddUUaCL8tXyySt8mJZqfpiE_SIentLjiupC__21DZwg",
			},
			want:    200,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMatchList(tt.args.region, tt.args.accountID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMatchList() error = %v, wantErr %v %v", err, tt.wantErr, ballotX)
				return
			}
			if !reflect.DeepEqual(got.StatusCode, tt.want) {
				t.Errorf("GetMatchList() = %v, want %v %v", got.StatusCode, tt.want, ballotX)
			}
			t.Logf("Checking Status Code 200 %v", checkMark)
		})
	}
}

func Test_makeRequest(t *testing.T) {
	type args struct {
		region   string
		endpoint string
		fn       request
	}
	tests := []struct {
		name    string
		args    args
		want    *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := makeRequest(tt.args.region, tt.args.endpoint, tt.args.fn)
			if (err != nil) != tt.wantErr {
				t.Errorf("makeRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMatchAPI(t *testing.T) {
	type args struct {
		matchID int
		region  string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Get Match information",
			args: args{
				matchID: 2856049818,
				region:  "na1",
			},
			want:    200,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMatchAPI(tt.args.matchID, tt.args.region)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMatchAPI() error = %v, wantErr %v %v", err, tt.wantErr, ballotX)
				return
			}
			if !reflect.DeepEqual(got.StatusCode, tt.want) {
				t.Errorf("GetMatchList() = %v, want %v %v", got.StatusCode, tt.want, ballotX)
			}
			t.Logf("Checking Status Code 200 %v", checkMark)
		})
	}
}
