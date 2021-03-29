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

// ConvertInt converts string value to int.
func ConvertInt(s string) (v int, err error) {
	var f float64
	f, err = ConvertFloat64(s)
	if err != nil {
		return
	}

	v = int(math.Round(f))

	return
}

// ConvertInt64 converts string value to int64.
func ConvertInt64(s string) (v int64, err error) {
	var f float64
	f, err = ConvertFloat64(s)
	if err != nil {
		return
	}

	v = int64(math.Round(f))

	return
}

// ConvertFloat64 converts string value to float64.
func ConvertFloat64(s string) (v float64, err error) {
	v, err = strconv.ParseFloat(sanitize(s), 64)
	return
}
