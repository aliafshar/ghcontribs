// Package ghcontribs counts all github commits for a user within the last
// twelve months.
// It relies on the github calendar that is displayed for each user as the API
// can't give this information.
package ghcontribs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Where the calendar gets its data.
const URI_TEMPLATE = "https://github.com/users/%v/contributions_calendar_data"

// fetchRawContribs gets the raw calendar data.
func fetchCalendar(username string) ([][]interface{}, error) {
	u := fmt.Sprintf(URI_TEMPLATE, username)
	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var v [][]interface{}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&v)
	if err != nil {
		return nil, err
	}
	return v, err
}

// parseDate parses the date format in github calendar info.
func parseDate(s string) (time.Time, error) {
	d, err := time.Parse("2006/01/02", s)
	if err != nil {
		return time.Time{}, err
	}
	return d, nil
}

// TotalContribsFor filters the user's calendar after a given date and totals
// the contributions.
func TotalContribsFor(username string, after time.Time) (float64, error) {
	cs, err := fetchCalendar(username)
	if err != nil {
		return 0, err
	}
	total := 0.0
	for _, c := range cs {
		s := c[0].(string)
		t := c[1].(float64)
		d, err := parseDate(s)
		if err != nil {
			return 0, err
		}
		if d.After(after) {
			total += t
		}
	}
	return total, nil
}
