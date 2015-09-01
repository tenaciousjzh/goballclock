package validator

import (
	"errors"
	"log"
	"strconv"
)

const min int = 27
const max int = 127
const InvalidBallInput string = "The first parameter must be a number between 27 and 127\n"
const InvalidDuration string = "The second parameter must be a number greater than 0\n"

//ParsedIntResult is used to indicate to the main entry point
//of the application the status after parsing program arguments
//IsValid: bool indicating if the arg is a valid one
//Value: the parsed integer value
//Error: if parsing failed, this will provide an error message
type ParsedIntResult struct {
	IsValid bool
	Value   int
	Error   error
}

//ValidateBallInput checks to see if the string argument
//provided can be parsed as an integer for the number of
//balls to put into the queue for the ball clock
func ValidateBallInput(input string) ParsedIntResult {
	numBalls, err := strconv.Atoi(input)
	isValid := true
	if err != nil {
		isValid = false
	}

	if IsValidBallRange(numBalls) == false {
		err = errors.New(InvalidBallInput)
		isValid = false
	}
	pr := ParsedIntResult{IsValid: isValid, Value: numBalls, Error: err}
	logIfInvalid(pr)
	return pr
}

//IsValidBallRange is used to determine if the argument provided
//to queue up for the number of balls is valid (between 27 and 127)
func IsValidBallRange(numBalls int) bool {
	if numBalls == 0 || numBalls < min || numBalls > max {
		return false
	}
	return true
}

//ValidateDuration checks if the second string argument provided
//is parsable as an integer greater than zero. This will put
//the ballclock application into its second mode to evaluate
//the duration provided and print out a json string
func ValidateDuration(input string) ParsedIntResult {
	duration, err := strconv.Atoi(input)
	isValid := true
	if err != nil {
		isValid = false
	}
	if IsValidDuration(duration) == false {
		isValid = false
		err = errors.New(InvalidDuration)
	}

	pr := ParsedIntResult{IsValid: isValid, Value: duration, Error: err}
	logIfInvalid(pr)
	return pr
}

//IsValidDuration checks to see if the duration provided is an
//integer value greater than or equal to zero to return true
func IsValidDuration(duration int) bool {
	if duration < 0 {
		return false
	}
	return true
}

func logIfInvalid(pr ParsedIntResult) {
	if pr.IsValid == false {
		go log.Printf(pr.Error.Error())
	}
}
