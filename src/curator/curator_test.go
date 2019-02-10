package curator

import (
	"fmt"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestCurator_Init(t *testing.T) {
	type fields struct {
		Splash map[string][]splashArt
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Curator{
				Splash: tt.fields.Splash,
			}
			c.Init()
		})
	}
}

func Test_loadSplashArt(t *testing.T) {
	t.Log("Given I read Champion Splash Arts Directory")
	{
		// var dir, err = os.Getwd()
		// if err != nil {
		// 	log.Fatal(err)
		// }
		splash := loadSplashArt("../" + championSplash)

		fmt.Println(splash)
	}
}

// func Test_loadSplashArt(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		want map[string][]splash
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name: "Get splash Art",
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			// if got := loadSplashArt(); !reflect.DeepEqual(got, tt.want) {
// 			// 	t.Errorf("loadSplashArt() = %v, want %v", got, tt.want)
// 			// }
// 			got := loadSplashArt()

// 			fmt.Println(got)
// 		})
// 	}
// }
