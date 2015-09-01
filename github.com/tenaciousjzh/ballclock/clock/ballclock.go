package clock

import (
	"encoding/json"
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
	stats        *clockStats
	showCounts   bool
}

type clockStats struct {
	halfDays     int
	fullDays     int
	minCount     int
	fiveMinCount int
	hourCount    int
}

func NewBallClock(numBalls int, duration int) (*ballClock, error) {

	if validator.IsValidBallRange(numBalls) == false {
		return nil, errors.New(validator.InvalidBallInput)
	}
	if validator.IsValidDuration(duration) == false {
		return nil, errors.New(validator.InvalidDuration)
	}
	bc := new(ballClock)
	bc.numBalls = numBalls
	bc.duration = duration
	bc.ballq = loadBallQueue(numBalls)
	bc.orig = make([]int, numBalls)
	copy(bc.orig, bc.ballq)
	bc.minTrack = make([]int, 0)
	bc.fiveMinTrack = make([]int, 0)
	bc.hourTrack = make([]int, 0)

	bc.reportMode = ReportDays
	bc.debug = false
	bc.showCounts = false
	stats := new(clockStats)
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
		return clock.simulateDuration()
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

type jsonResult struct {
	Min     []int
	FiveMin []int
	Hour    []int
	Main    []int
}

func (clock *ballClock) simulateDuration() string {
	ballVal := -1
	for i := 0; i < clock.duration; i++ {
		ballVal, clock.ballq = slice.Shift(clock.ballq)
		clock.updMinTrack(ballVal)
	}
	durationStats := jsonResult{
		Min:     clock.minTrack,
		FiveMin: clock.fiveMinTrack,
		Hour:    clock.hourTrack,
		Main:    clock.ballq,
	}

	result, err := json.Marshal(durationStats)
	if err != nil {
		log.Println("error: ", err)
	}
	return string(result)
}

func (clock *ballClock) updMinTrack(ballVal int) {
	if clock.debug {
		log.Println("==========================================")
		log.Println("Calling updMinTrack")
	}
	stats := clock.stats
	stats.minCount += 1
	if len(clock.minTrack) == minTrackCapacity {
		if clock.debug {
			log.Printf("minTrack full! Adding balVall %d to next track.", ballVal)
			log.Printf("minTrack length: %d, minTrackCapacity: %d", len(clock.minTrack), minTrackCapacity)
		}
		//Put the balls on the track in reverse order
		//back into the ballq
		for i, val := range clock.minTrack {
			if clock.debug && clock.showCounts {
				log.Printf("minTrack[%d] = %d", i, val)
			}
			lastVal := 0
			lastVal, clock.minTrack = slice.Pop(clock.minTrack)
			clock.ballq = slice.Push(clock.ballq, lastVal)
		}
		//Put the next ball from ballq into the next track
		clock.updFiveMinTrack(ballVal)
	} else {
		if clock.debug {
			log.Printf("adding ballVal %d to minTrack", ballVal)
		}
		clock.minTrack = slice.Push(clock.minTrack, ballVal)

	}
	if clock.debug {
		clock.PrintDiagnostic()
	}
}

func (clock *ballClock) updFiveMinTrack(ballVal int) {
	if clock.debug {
		log.Println("==========================================")
		log.Println("Calling updFiveMinTrack")
	}
	stats := clock.stats
	stats.fiveMinCount += 1
	if len(clock.fiveMinTrack) == fiveMinTrackCapacity {
		if clock.debug {
			log.Printf("fiveMinTrack full! Adding balVall %d to next track.", ballVal)
			log.Printf("fiveMinTrack length: %d, fiveMinTrackCapacity: %d", len(clock.fiveMinTrack), fiveMinTrackCapacity)
		}

		for i, val := range clock.fiveMinTrack {
			if clock.debug && clock.showCounts {
				log.Printf("fiveMinTrack[%d] = %d", i, val)
			}
			lastVal := 0
			lastVal, clock.fiveMinTrack = slice.Pop(clock.fiveMinTrack)
			clock.ballq = slice.Push(clock.ballq, lastVal)
		}

		clock.updHourTrack(ballVal)
	} else {
		if clock.debug {
			log.Printf("adding ballVal %d to fiveMinTrack", ballVal)
		}
		clock.fiveMinTrack = slice.Push(clock.fiveMinTrack, ballVal)
	}
}

func (clock *ballClock) updHourTrack(ballVal int) {
	if clock.debug {
		log.Println("==========================================")
		log.Println("Calling updHourTrack")
	}
	stats := clock.stats
	stats.hourCount += 1
	if len(clock.hourTrack) == hourTrackCapacity {
		if clock.debug {
			log.Printf("hourTrack full! Adding balVall %d to next track.", ballVal)
			log.Printf("hourTrack length: %d, hourTrackCapacity: %d", len(clock.hourTrack), hourTrackCapacity)
		}
		for i, val := range clock.hourTrack {
			if clock.debug && clock.showCounts {
				log.Printf("hourTrack[%d] = %d", i, val)
			}
			lastVal := 0
			lastVal, clock.hourTrack = slice.Pop(clock.hourTrack)
			clock.ballq = slice.Push(clock.ballq, lastVal)
		}

		clock.ballq = slice.Push(clock.ballq, ballVal)
		stats.halfDays += 1
		clock.printCounts(ballVal)
	} else {
		if clock.debug {
			log.Printf("adding ballVal %d to hourTrack", ballVal)
		}
		clock.hourTrack = slice.Push(clock.hourTrack, ballVal)
	}
}

func (clock *ballClock) printCounts(ballVal int) {
	if clock.debug {
		log.Println("====================================")
		log.Println("============== Stats ===============")
		s := clock.stats
		log.Printf("ballVal: %d\n", ballVal)
		log.Printf("minCount: %d, fiveMinCount: %d, hourCount: %d\n", s.minCount, s.fiveMinCount, s.hourCount)
		log.Printf("halfDays: %d", s.halfDays)
		log.Println("====================================")
	}
}

func (clock *ballClock) PrintDiagnostic() {
	if clock.debug {
		log.Printf("ballq: %s\n", clock.ballq)
		log.Printf("minTrack: %s\n", clock.minTrack)
		log.Printf("fiveMinTrack: %s\n", clock.fiveMinTrack)
		log.Printf("hourTrack: %s\n", clock.hourTrack)
	}
}
