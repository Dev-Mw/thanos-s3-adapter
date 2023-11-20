package utils

import (
	"errors"
	"fmt"
	"time"
)

type ArgumentTypeError struct{ Err error }

func (m *ArgumentTypeError) Error() string {
	return fmt.Sprintf("ArgumentTypeError: %v", m.Err)
}

func DatesValidation(startDateStr string, endDateStr string, interval int) ([][]interface{}, error) {
	var ranges [][]interface{}

	startTimeUnix, err := time.Parse(time.RFC3339, startDateStr)
	if err != nil {
		return make([][]interface{}, 0), &ArgumentTypeError{Err: errors.New(err.Error())}
	}
	endTimeUnix, err := time.Parse(time.RFC3339, endDateStr)
	if err != nil {
		return make([][]interface{}, 0), &ArgumentTypeError{Err: errors.New(err.Error())}
	}
	maxDate := time.Now().UTC().Add(time.Hour * -1)
	intervalDuration := time.Duration(interval) * time.Second

	if startTimeUnix.After(endTimeUnix) {
		return make([][]interface{}, 0), &ArgumentTypeError{Err: errors.New("the StartDate cannot be greater than the EndDate")}
	}
	if startTimeUnix == endTimeUnix {
		return make([][]interface{}, 0), &ArgumentTypeError{Err: errors.New("the StartDate cannot be equal to the EndDate")}
	}
	if startTimeUnix.After(maxDate) {
		msg := "the StartDate allowed cannot be greater than " + maxDate.Format(time.RFC3339)
		return make([][]interface{}, 0), &ArgumentTypeError{Err: errors.New(msg)}
	}

	curr := startTimeUnix
	for curr.Before(endTimeUnix) {
		dates := []interface{}{curr.Format(time.RFC3339), curr.Add(intervalDuration - time.Millisecond).Format(time.RFC3339)}
		ranges = append(ranges, dates)
		curr = curr.Add(intervalDuration)
	}

	return ranges, nil
}
