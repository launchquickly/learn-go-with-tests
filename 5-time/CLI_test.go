package poker_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	poker "github.com/launchquickly/learn-go-with-tests/5-time"
)

var dummyBlindAlerter = &poker.SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}
var dummyStdout = &bytes.Buffer{}

func TestCLI(t *testing.T) {
	t.Run("start game with 3 players and record 'Chris' as winner", func(t *testing.T) {
		game := &GameSpy{}
		stdout := &bytes.Buffer{}

		in := userSends("3", "Chris")
		cli := poker.NewCLI(in, stdout, game)

		cli.PlayPoker()

		assertMessageSentToUser(t, stdout, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 3)
		assertFinishCalledWith(t, game, "Chris")
	})

	t.Run("start game with 8 players and record 'Cleo' as winner", func(t *testing.T) {
		game := &GameSpy{}

		in := userSends("8", "Cleo")
		cli := poker.NewCLI(in, dummyStdout, game)

		cli.PlayPoker()

		assertGameStartedWith(t, game, 8)
		assertFinishCalledWith(t, game, "Cleo")
	})

	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		game := &GameSpy{}

		stdout := &bytes.Buffer{}
		in := userSends("Pies")

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertGameNotStarted(t, game)
		assertMessageSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})
}

type GameSpy struct {
	StartCalled  bool
	StartedWith  int
	FinishedWith string
}

func (g *GameSpy) Start(numberOfPlayers int) {
	g.StartCalled = true
	g.StartedWith = numberOfPlayers
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedWith = winner
}

func assertFinishCalledWith(t testing.TB, game *GameSpy, winner string) {
	t.Helper()
	if game.FinishedWith != winner {
		t.Errorf("did not finish with correct winner got %q want %q", game.FinishedWith, winner)
	}
}

func assertGameNotStarted(t testing.TB, game *GameSpy) {
	t.Helper()
	if game.StartCalled {
		t.Error("game should not have started")
	}
}

func assertGameStartedWith(t testing.TB, game *GameSpy, noOfPlayers int) {
	t.Helper()
	if game.StartedWith != noOfPlayers {
		t.Errorf("wanted Start called with %d but got %d", noOfPlayers, game.StartedWith)
	}
}

func assertMessageSentToUser(t testing.TB, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	got := stdout.String()
	want := strings.Join(messages, "")

	if got != want {
		t.Errorf("got %q, sent to stdout but expected %+v", got, want)
	}
}

func userSends(input ...string) io.Reader {
	return strings.NewReader(strings.Join(input, "\n"))
}
