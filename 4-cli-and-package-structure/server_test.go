package poker_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	poker "github.com/launchquickly/learn-go-with-tests/4-cli-and-package-structure"
)

func TestGETPlayers(t *testing.T) {
	store := poker.NewStubPlayerStore(
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		nil,
		nil,
	)
	server := poker.NewPlayerServer(store)

	tests := []struct {
		name               string
		player             string
		expectedHTTPStatus int
		expectedScore      string
	}{
		{
			name:               "returns Pepper's score",
			player:             "Pepper",
			expectedHTTPStatus: http.StatusOK,
			expectedScore:      "20",
		},
		{
			name:               "returns Floyd's score",
			player:             "Floyd",
			expectedHTTPStatus: http.StatusOK,
			expectedScore:      "10",
		},
		{
			name:               "returns 404 on missing player",
			player:             "Apollo",
			expectedHTTPStatus: http.StatusNotFound,
			expectedScore:      "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := poker.TestNewGetScoreRequest(tt.player)
			response := httptest.NewRecorder()

			server.ServeHTTP(response, request)

			poker.AssertStatus(t, response.Code, tt.expectedHTTPStatus)
			poker.AssertResponseBody(t, response.Body.String(), tt.expectedScore)
		})
	}
}

func TestLeague(t *testing.T) {
	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wantedLeague := poker.League{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}

		store := poker.NewStubPlayerStore(nil, nil, wantedLeague)
		server := poker.NewPlayerServer(store)

		request := poker.TestNewLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := poker.TestGetLeagueFromResponse(t, response.Body)
		poker.AssertStatus(t, response.Code, http.StatusOK)
		poker.AssertLeague(t, got, wantedLeague)
		poker.AssertContentType(t, response, poker.JsonContentType)
	})
}

func TestStoreWins(t *testing.T) {
	store := poker.NewStubPlayerStore(
		map[string]int{},
		nil,
		nil,
	)
	server := poker.NewPlayerServer(store)

	t.Run("it records wins when POST", func(t *testing.T) {
		player := "Pepper"

		request := poker.TestNewPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertStatus(t, response.Code, http.StatusAccepted)

		poker.AssertPlayerWin(t, store, player)
	})
}
