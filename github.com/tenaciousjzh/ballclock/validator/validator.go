package validator

import (
	"errors"
	"log"
	"strconv"
)

const min int = 27
const max int = 127
const invalidBallInput string = "The first parameter must be a number between 27 and 127\n"
const invalidDuration string = "The second parameter must be a number greater than 0\n"

type ParsedIntResult struct {
	IsValid bool
	Value   int
	Error   error
}

func ValidateBallInput(input string) ParsedIntResult {
	numBalls, err := strconv.Atoi(input)
	isValid := true
	if err != nil {
		isValid = false
	}
	if numBalls == 0 || numBalls < min || numBalls > max {
		err = errors.New(invalidBallInput)
		isValid = false
	}
	pr := ParsedIntResult{IsValid: isValid, Value: numBalls, Error: err}
	logIfInvalid(pr)
	return pr
}

func ValidateDuration(input string) ParsedIntResult {
	duration, err := strconv.Atoi(input)
	isValid := true
	if err != nil {
		isValid = false
	}
	if duration < 0 {
		isValid = false
		err = errors.New(invalidDuration)
	}

	pr := ParsedIntResult{IsValid: isValid, Value: duration, Error: err}
	logIfInvalid(pr)
	return pr
}

func logIfInvalid(pr ParsedIntResult) {
	if pr.IsValid == false {
		go log.Printf(pr.Error.Error())
	}
}
