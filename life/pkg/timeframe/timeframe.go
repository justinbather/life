package timeframe

import (
	"errors"
	"strings"
	"time"
)

var ERR_INVALID_TIMEFRAME error = errors.New("Invalid timeframe.")

type timeframe int

func (t timeframe) String() string {
	return timeframeStrings[t]
}

func (t timeframe) GetRange() map[string]string {

	now := time.Now()
	eod := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location()).Format(time.RFC3339)

	switch t {
	case TODAY:
		start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).Format(time.RFC3339)

		return map[string]string{"start": start, "end": eod}
	case WEEK:
		sevenDaysAgo := now.AddDate(0, 0, -7).Format(time.RFC3339)
		return map[string]string{"start": sevenDaysAgo, "end": eod}
	case MONTH:
		monthAgo := now.AddDate(0, -1, 0).Format(time.RFC3339)
		return map[string]string{"start": monthAgo, "end": eod}
	case YEAR:
		yearAgo := now.AddDate(-1, 0, 0).Format(time.RFC3339)
		return map[string]string{"start": yearAgo, "end": eod}
	}

	return nil
}

const (
	UNKNOWN timeframe = iota
	TODAY
	WEEK
	MONTH
	YEAR
)

var timeframeStrings = map[timeframe]string{
	UNKNOWN: "unknown",
	TODAY:   "today",
	WEEK:    "week",
	MONTH:   "month",
	YEAR:    "year",
}

func ParseTimeframe(val string) (timeframe, error) {
	switch strings.ToLower(val) {
	case "today":
		return TODAY, nil
	case "week":
		return WEEK, nil
	case "month":
		return MONTH, nil
	case "year":
		return YEAR, nil
	default:
		return UNKNOWN, ERR_INVALID_TIMEFRAME
	}
}
