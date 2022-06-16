package poker_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	poker "github.com/launchquickly/learn-go-with-tests/6-websockets"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabse := poker.TestCreateTempFile(t, "")
	defer cleanDatabse()

	store, err := poker.NewFileSystemPlayerStore(database)
	poker.AssertNoError(t, err)

	server := mustMakePlayerServer(t, store, dummyGame)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), poker.TestNewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), poker.TestNewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), poker.TestNewPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, poker.TestNewGetScoreRequest(player))
		poker.AssertStatus(t, response, http.StatusOK)

		poker.AssertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, poker.TestNewLeagueRequest())
		poker.AssertStatus(t, response, http.StatusOK)

		got := poker.TestGetLeagueFromResponse(t, response.Body)
		want := poker.League{
			{"Pepper", 3},
		}
		poker.AssertLeague(t, got, want)
	})
}
