package curator

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

/**
This package "curate" the static content.
Reading and parse the images and contents into usable map
*/

// Init initialize Curator instance to gather static data information including the file/image path
func (c *Curator) Init() {
	c.Splash = loadSplashArt(championSplash)
}

func loadSplashArt(path string) map[string][]splashArt {

	m := make(map[string][]splashArt)

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		s := strings.FieldsFunc(f.Name(), func(r rune) bool {
			return r == '_' || r == '.'
		})
		champ := s[0]
		index, _ := strconv.Atoi(s[1])
		// fmt.Println(s, champ, index, f.Name())

		m[champ] = append(m[champ], splashArt{
			ChampionName: champ,
			ID:           index,
			FileName:     f.Name(),
			URL:          championSplash + f.Name(),
		})
	}

	return m
}
