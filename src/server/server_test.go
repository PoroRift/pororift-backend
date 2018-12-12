package main

// import (
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/PoroRift/pororift-backend/lol"
// 	"github.com/labstack/echo"
// 	"github.com/stretchr/testify/assert"
// )

// // "net/http"
// // // "net/http/httptest"

// // func Test(t *testting) {
// // 	// e := echo.New()
// // 	// req := httptest.NewRequest(http.MethodPost, "/", nil)
// // 	// rec := httptest.NewRecorder()
// // 	// c := e.NewContext(req, rec)
// // }

// var expected_stat = &lol.Stat{
// 	Id: "_fFpQyv9Ml4bqMequtgYFuibmGIbmTqUEN3lmu29-DxdN8c",
// }

// func TestGetSummonerStat(t *testing.T) {

// 	// Setup
// 	e := echo.New()
// 	req := httptest.NewRequest(http.MethodGet, "/", nil)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	c.SetPath("/summoner/:name")
// 	c.SetParamNames("name")
// 	c.SetParamValues("richerthanu")

// 	// Assertions
// 	// if assert.NoError(t, h.)
// 	if assert.NoError(t, getSummonerStat(c)) {
// 		assert.Equal(t, http.StatusOK, rec.Code)
// 		var stat = &lol.Stat{}
// 		json.NewDecoder(rec.Body).Decode(&stat)
// 		// assert.Equal(t, *expected_stat, rec.Body.String())
// 		assert.Equal(t, expected_stat.Id, stat.Id)

// 	}

// }
