package main

import "testing"

func testDateParsin(t *testing.T) {
	date1 := "12:00:12"
	date2 := "2006-01-02 15:04:05"
	date3 := "2006-01-02 15:04:05"

	dates := []string{date1, date2, date3}

	for _, v := range dates {
		_, err := parseDate(v)

		if err != nil {
			t.Errorf("Got error %s", err)
		}
	}

}
