package services

import {
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
}

type Match struct {
	GameId int64 `json:"gameId"`
	PlatformId string `json:"platformId"`
	GameCreation int64 `json:"gameCreation"`
	GameDuration int64 `json:"gameDuration"`
	QueueId int `json:"queueId"`
	MapId int `json:"mapId"`
	SeasonId int `json:"seasonId"`
	GameVersion string `json:"gameVersion"`
	GameMode string `json:"gameMode"`
	GameType string `json:"gameType"`
}


func init() {
	d
}

