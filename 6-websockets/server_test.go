package poker_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	poker "github.com/launchquickly/learn-go-with-tests/6-websockets"
)

var dummyGame = &GameSpy{}

func TestGame(t *testing.T) {

	t.Run("game request returns OK status", func(t *testing.T) {
		server, _ := poker.NewPlayerServer(&poker.StubPlayerStore{}, dummyGame)

		request := poker.TestNewGameRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertStatus(t, response, http.StatusOK)
	})

	t.Run("start a game with 3 players and declare Ruth the winner", func(t *testing.T) {
		wantedBlindAlert := "Blind is 300"
		winner := "Ruth"

		game := &GameSpy{BlindAlert: []byte(wantedBlindAlert)}
		server := httptest.NewServer(mustMakePlayerServer(t, dummyPlayerStore, game))
		ws := mustDialWS(t, "ws"+strings.TrimPrefix(server.URL, "http")+"/ws")

		defer server.Close()
		defer ws.Close()

		writeWSMessage(t, ws, "3")
		writeWSMessage(t, ws, winner)

		assertGameStartedWith(t, game, 3)
		assertFinishCalledWith(t, game, "Ruth")
		within(t, 10*time.Millisecond, func() { assertWebsocketGotMsg(t, ws, wantedBlindAlert) })
	})
}

func TestGETPlayers(t *testing.T) {
	store := poker.NewStubPlayerStore(
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		nil,
		nil,
	)
	server := mustMakePlayerServer(t, store, dummyGame)

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

			poker.AssertStatus(t, response, tt.expectedHTTPStatus)
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
		server := mustMakePlayerServer(t, store, dummyGame)

		request := poker.TestNewLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := poker.TestGetLeagueFromResponse(t, response.Body)
		poker.AssertStatus(t, response, http.StatusOK)
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
	server := mustMakePlayerServer(t, store, dummyGame)

	t.Run("it records wins when POST", func(t *testing.T) {
		player := "Pepper"

		request := poker.TestNewPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertStatus(t, response, http.StatusAccepted)

		poker.AssertPlayerWin(t, store, player)
	})
}

func assertWebsocketGotMsg(t *testing.T, ws *websocket.Conn, want string) {
	_, msg, _ := ws.ReadMessage()
	if string(msg) != want {
		t.Errorf(`got "%s", want "%s"`, string(msg), want)
	}
}

func mustDialWS(t *testing.T, url string) *websocket.Conn {
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("could not open a ws connection on %s %v", url, err)
	}
	return ws
}

func mustMakePlayerServer(t *testing.T, store poker.PlayerStore, game poker.Game) *poker.PlayerServer {
	server, err := poker.NewPlayerServer(store, game)
	if err != nil {
		t.Fatal("problem creating player server", err)
	}
	return server
}

func within(t testing.TB, d time.Duration, assert func()) {
	t.Helper()

	done := make(chan struct{}, 1)

	go func() {
		assert()
		done <- struct{}{}
	}()

	select {
	case <-time.After(d):
		t.Error("timed out")
	case <-done:
	}
}

func writeWSMessage(t testing.TB, conn *websocket.Conn, message string) {
	t.Helper()
	if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		t.Fatalf("could not send connection over we connection %v", err)
	}
}
