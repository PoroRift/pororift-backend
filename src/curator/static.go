package curator

/**
Static will contain all static path to the files and configuration.
*/
import (
	"fmt"
)

var staticPath = "/static"
var imagePath = staticPath + "/img"
var championSplash = fmt.Sprintf("%s/champion/splash", imagePath)
var championloading = imagePath + "/champion/loading"
var championTiles = imagePath + "/champion/tiles"

// Init function will read the static data content
func Init() {

	// cp := loadSplash()

	// fmt.Println(cp)
}

// func loadSplash() map[string][]splash {

// 	files, err := ioutil.ReadDir(championSplash)
// 	m := make(map[string][]splash)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for _, f := range files {
// 		// fmt.Println(f.Name())
// 		// s := strings.Split(f.Name(), "_")
// 		s := strings.FieldsFunc(f.Name(), func(r rune) bool {
// 			return r == '_' || r == '.'
// 		})

// 		champ := s[0]
// 		index, _ := strconv.Atoi(s[1])
// 		fmt.Println(champ, index, f.Name())

// 		// // m[champ] =
// 		// if val, ok := m[champ]; !ok {
// 		// 	var s []splash
// 		// 	m[champ] = []splash
// 		// }

// 		m[champ] = append(m[champ], splash{
// 			ChampName: champ,
// 			ID:        index,
// 			FileName:  f.Name(),
// 			URL:       "static/img/champion/splash/" + f.Name(),
// 		})
// 	}

// 	fmt.Println(m)
// 	return nil
// }

type (
	// Curator contain information about the static content (not Riot API related) such as champion, art, item, etc.
	Curator struct {
		Splash map[string][]splashArt `json:"splashArt"`
	}

	splashArt struct {
		ChampionName string `json:"champ"`
		ID           int    `json:"id"`
		FileName     string `json:"fileName"`
		URL          string `json:"url"`
	}
)

// type (
// 	splash struct {
// 		ChampName string `json:"champ"`
// 		ID        int    `json:"id"`
// 		FileName  string `json:"fileName"`
// 		URL       string `json:"url"`
// 	}
// )
