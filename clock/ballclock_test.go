package clock

import (
	"github.com/x/goballclock/validator"
	"testing"
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

func TestNewBallClockWithAboveMaxBallCount(t *testing.T) {
	numBalls, duration := 128, 0
	bc, err := NewBallClock(numBalls, duration)
	if bc != nil {
		t.Errorf("Ballclock should have been set to nil based on invalid argument of %d provided.", numBalls)
	}
	if err == nil {
		t.Errorf("With numBalls value of %d, error should have been: %s", numBalls, validator.InvalidBallInput)
	}
}

func TestNewBallClockValidBallCount(t *testing.T) {
	numBalls, duration := 30, 0
	bc, err := NewBallClock(numBalls, duration)
	if bc == nil || err != nil {
		t.Errorf("Ballclock should have been initialized correctly with numBalls value of: %d, and duration: %d", numBalls, duration)
	}
}
