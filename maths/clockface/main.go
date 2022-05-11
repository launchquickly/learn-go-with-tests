// Writes an SVG clockface of the current time to Stdout.
package main

import (
	"os"
	"time"

	"github.com/launchquickly/learn-go-with-tests/maths/svg"
)

func main() {
	t := time.Now()
	svg.Write(os.Stdout, t)
}
