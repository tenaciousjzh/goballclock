package clock

import (
	"errors"
	"github.com/tenaciousjzh/ballclock/util"
	"github.com/tenaciousjzh/ballclock/validator"
	"log"
	"strconv"
)

const minTrackCapacity int = 4
const fiveMinTrackCapacity int = 11
const hourTrackCapacity = 11

type mode int

const (
	ReportDays = iota
	ReportTracks
)

type ballClock struct {
	numBalls     int
	duration     int
	ballq        []int
	orig         []int
	minTrack     []int
	fiveMinTrack []int
	hourTrack    []int
	reportMode   mode
	debug        bool
	stats        clockStats
}

type clockStats struct {
	halfDays     int
	fullDays     int
	minCount     int
	fiveMinCount int
	hourCount    int
}

func NewBallClock(numBalls int, duration int) (*ballClock, error) {
	bc := new(ballClock)
	bc.numBalls = numBalls
	if validator.IsValidBallRange(numBalls) == false {
		return bc, errors.New(validator.InvalidBallInput)
	}
	if validator.IsValidDuration(duration) == false {
		return bc, errors.New(validator.InvalidDuration)
	}

	bc.ballq = loadBallQueue(numBalls)
	bc.orig = make([]int, numBalls)
	copy(bc.orig, bc.ballq)
	bc.minTrack = make([]int, minTrackCapacity)
	bc.fiveMinTrack = make([]int, fiveMinTrackCapacity)
	bc.hourTrack = make([]int, hourTrackCapacity)

	bc.reportMode = ReportDays
	bc.debug = false
	stats := clockStats{halfDays: 0, fullDays: 0, minCount: 0, fiveMinCount: 0, hourCount: 0}
	bc.stats = stats

	if duration > 0 {
		bc.reportMode = ReportTracks
		bc.duration = duration
	}
	return bc, nil
}

func loadBallQueue(size int) []int {
	q := make([]int, size)
	for i := 0; i < len(q); i++ {
		q[i] = i + 1
	}
	return q
}

func (clock *ballClock) RunClock() string {
	if clock.reportMode == ReportDays {
		return clock.simulateDays()
	} else {
		//return clock.simulateDuration()
		return ""
	}
}

func (clock *ballClock) simulateDays() string {
	stats := clock.stats
	stats.halfDays = 0
	counter := 0
	ballVal := -1
	for {
		ballVal, clock.ballq = slice.Shift(clock.ballq)
		clock.updMinTrack(ballVal)
		counter++
		if slice.Same(clock.ballq, clock.orig) {
			break
		}
	}
	stats.fullDays = stats.halfDays / 2
	return strconv.Itoa(clock.numBalls) + " balls cycle for " + strconv.Itoa(stats.fullDays) + " days."

}

func (clock *ballClock) updMinTrack(ballVal int) {
	stats := clock.stats
	stats.minCount += 1
	if len(clock.minTrack) == minTrackCapacity {
		//Put the balles on the track in reverse order
		//back into the ballq
		for i, val := range clock.minTrack {
			lastVal := 0
			lastVal, clock.minTrack = slice.Pop(clock.minTrack)
			clock.ballq = slice.Push(clock.ballq, lastVal)
		}
		//Put the next ball from ballq into the next track
		clock.updFiveMinTrack(ballVal)
	} else {
		clock.minTrack = slice.Push(clock.minTrack, ballVal)
	}
}

func (clock *ballClock) updFiveMinTrack(ballVal int) {
	stats := clock.stats
	stats.fiveMinCount += 1
	if len(clock.fiveMinTrack) == fiveMinTrackCapacity {
		for i, v := range clock.fiveMinTrack {
			lastVal := 0
			lastVal, clock.fiveMinTrack = slice.Pop(clock.fiveMinTrack)
			clock.ballq = slice.Push(clock.ballq, lastVal)
		}
		clock.updHourTrack(ballVal)
	} else {
		clock.fiveMinTrack = slice.Push(clock.fiveMinTrack, ballVal)
	}
}

func (clock *ballClock) updHourTrack(ballVal int) {
	stats := clock.stats
	stats.hourCount += 1
	if len(clock.hourTrack) == hourTrackCapacity {
		for i, val := range clock.hourTrack {
			lastVal := 0
			lastVal, clock.hourTrack = slice.Pop(clock.hourTrack)
			clock.ballq = slice.Push(clock.ballq, lastVal)
		}
		clock.ballq = slice.Push(clock.ballq, ballVal)
		stats.halfDays += 1
		clock.printCounts(ballVal)
	} else {
		clock.hourTrack = slice.Push(clock.hourTrack, ballVal)
	}
}

func (clock *ballClock) printCounts(ballVal int) {
	if clock.debug {
		log.Printf("ballVal: %d\n", ballVal)
		log.Printf("minCount: %d, fiveMinCount: %d, hourCount: %d\n", clock.stats.minCount, clock.stats.fiveMinCount, clock.stats.hourCount)
		log.Printf("halfDays: %d", clock.stats.halfDays)
	}
}
