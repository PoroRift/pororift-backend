// Package data provide and compile all data that the handler(from client) requested
package data

import "fmt"

var db = &DataBase{}

func init() {
	db.Test = "Test String"
}

// Look up player by name
// Return player information
func GetPlayer(name string) {
	fmt.Println(name)
	fmt.Println(db.Test)
}
