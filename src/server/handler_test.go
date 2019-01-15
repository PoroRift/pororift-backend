package main

import (
	"testing"
)

func Test_getMatch(t *testing.T) {
	tests := []struct {
		name    string
		args    int
		wantErr bool
	}{
		{
			name:    "GET Match: 2856049818",
			args:    2856049818,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if err := getMatch(tt.args.c); (err != nil) != tt.wantErr {
			// 	t.Errorf("getMatch() error = %v, wantErr %v", err, tt.wantErr)
			// }
		})
	}
}
