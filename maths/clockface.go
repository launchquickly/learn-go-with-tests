// Package clockface provides functions that calculate the positions of the hands.
// of an analogue clock.
package clockface

import (
	"math"
	"time"
)

const (
	hoursInHalfClock   = 6
	hoursInClock       = 2 * 6
	minutesInHalfClock = 30
	minutesInClock     = 2 * 30
	secondsInHalfClock = 30
)

// A Point is a Cartesian coordinate. They are used in the package.
// to represent the unit vector from the origin of a clock hand.
type Point struct {
	X float64
	Y float64
}

// HourHandPoint is the unit vector of the hour hand at time `t`,
// represented by a Point.
func HourHandPoint(t time.Time) Point {
	return angleToPoint(HoursInRadians(t))
}

func HoursInRadians(t time.Time) float64 {
	return (MinutesInRadians(t) / hoursInClock) +
		(math.Pi / (hoursInHalfClock / float64(t.Hour()%hoursInClock)))
}

// MinuteHandPoint is the unit vector of the minute hand at time `t`,
// represented by a Point.
func MinuteHandPoint(t time.Time) Point {
	return angleToPoint(MinutesInRadians(t))
}

func MinutesInRadians(t time.Time) float64 {
	return (SecondsInRadians(t) / minutesInClock) +
		(math.Pi / (minutesInHalfClock / float64(t.Minute())))
}

// SecondHandPoint is the unit vector of the minute hand at time `t`,
// represented by a Point.
func SecondHandPoint(t time.Time) Point {
	return angleToPoint(SecondsInRadians(t))
}

func SecondsInRadians(t time.Time) float64 {
	return math.Pi / (secondsInHalfClock / float64(t.Second()))
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}
