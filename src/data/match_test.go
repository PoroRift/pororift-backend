package data

import (
	"sync"
	"testing"
)

func TestMatchInit(t *testing.T) {
	type fields struct {
		GameID int
		mutex  *sync.Mutex
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Get Match Information",
			fields: fields{
				GameID: 2899743601,
				mutex:  &sync.Mutex{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &match{
				GameID: tt.fields.GameID,
				mutex:  tt.fields.mutex,
			}
			if err := m.Init(); (err != nil) != tt.wantErr {
				t.Errorf("Match.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
			// t.Log(m.Data)
		})
	}
}
