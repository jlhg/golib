package strconv

import (
	"math"
	"regexp"
	"strconv"
)

var (
	sanitizePattern = regexp.MustCompile(`[^0-9.]`)
)

func sanitize(s string) string {
	return sanitizePattern.ReplaceAllString(s, "")
}

// ParseInt parses string value to int.
func ParseInt(s string) (v int, err error) {
	var f float64
	f, err = ParseFloat64(s)
	if err != nil {
		return
	}

	v = int(math.Round(f))

	return
}

// ParseInt64 parses string value to int64.
func ParseInt64(s string) (v int64, err error) {
	var f float64
	f, err = ParseFloat64(s)
	if err != nil {
		return
	}

	v = int64(math.Round(f))

	return
}

// ParseFloat64 parses string value to float64.
func ParseFloat64(s string) (v float64, err error) {
	v, err = strconv.ParseFloat(sanitize(s), 64)
	return
}
