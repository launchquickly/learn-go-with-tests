package poker_test

import (
	"bytes"
	"io"
	"strings"
	"testing"
	"time"

	poker "github.com/launchquickly/learn-go-with-tests/6-websockets"
)

var dummyBlindAlerter = &poker.SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}
var dummyOut = &bytes.Buffer{}

func TestCLI(t *testing.T) {
	t.Run("start game with 3 players and record 'Chris' as winner", func(t *testing.T) {
		game := &GameSpy{}
		out := &bytes.Buffer{}

		in := userSends("3", "Chris")
		cli := poker.NewCLI(in, out, game)

		cli.PlayPoker()

		assertMessageSentToUser(t, out, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 3)
		assertFinishCalledWith(t, game, "Chris")
	})

	t.Run("start game with 8 players and record 'Cleo' as winner", func(t *testing.T) {
		game := &GameSpy{}

		in := userSends("8", "Cleo")
		cli := poker.NewCLI(in, dummyOut, game)

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
	BlindAlert   []byte
	FinishedWith string
}

func (g *GameSpy) Start(numberOfPlayers int, out io.Writer) {
	g.StartCalled = true
	g.StartedWith = numberOfPlayers
	out.Write(g.BlindAlert)
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedWith = winner
}

func assertFinishCalledWith(t testing.TB, game *GameSpy, winner string) {
	t.Helper()

	passed := retryUntil(500*time.Millisecond, func() bool {
		return game.FinishedWith == winner
	})

	if !passed {
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

	passed := retryUntil(500*time.Millisecond, func() bool {
		return game.StartedWith == noOfPlayers
	})

	if !passed {
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

func retryUntil(d time.Duration, f func() bool) bool {
	deadline := time.Now().Add(d)
	for time.Now().Before(deadline) {
		if f() {
			return true
		}
	}
	return false
}

func userSends(input ...string) io.Reader {
	return strings.NewReader(strings.Join(input, "\n"))
}
