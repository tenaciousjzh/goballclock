package clock

import (
	"testing"
	"validator"
)

func TestNewBallClockBelowMinBallCount(t *testing.T) {
	numBalls, duration := 26, 0
	bc, err := NewBallClock(numBalls, duration)
	if bc != nil {
		t.Errorf("Ballclock should have been set to nil based on invalid argument of %d provided.", numBalls)
	}
	if err == nil {
		t.Errorf("With numBalls value of %d, error should have been: %s", numBalls, validator.InvalidBallInput)
	}
}
