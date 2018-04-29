package main

import (
	"errors"
	"time"
)

func tryDateFormats(userTimeValue string, dateFormats []string) (*time.Time, error) {
	var localLocation = time.Now().Local().Location()

	for _, layout := range dateFormats {
		t, err := time.ParseInLocation(layout, userTimeValue, localLocation)

		if err == nil {
			return &t, nil
		}
	}

	return nil, errors.New("Unable to parse date: " + userTimeValue)
}

func parseDate(userTimeValue string) (*time.Time, error) {
	var layouts = []string{"2006-01-02 15:04:05", "2006-01-02 15:04", "2006-01-02"}

	parsedValue, err := tryDateFormats(userTimeValue, layouts)

	if err == nil {
		return parsedValue, nil
	}

	// Try parsing with date string
	var currentTime = time.Now().Local()
	var datePrefix = currentTime.Format("2006-01-02")

	var userTimeValueWithPrefix = datePrefix + " " + userTimeValue

	parsedValue, err = tryDateFormats(userTimeValueWithPrefix, layouts)

	// Unable to parse, exit
	if err != nil {
		return nil, err
	}

	// Got hour value in the past (assume tomorrow is meant)
	if parsedValue.Before(currentTime) {
		var tomorrow = currentTime.AddDate(0, 0, 1)
		var datePrefix = tomorrow.Format("2006-01-02")
		var userTimeValueWithPrefixTomorrow = datePrefix + " " + userTimeValue
		return tryDateFormats(userTimeValueWithPrefixTomorrow, layouts)
	}

	return parsedValue, nil
}
