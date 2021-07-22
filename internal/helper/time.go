package helper

import (
	"errors"
	"time"
)

func GetCurrentMilis() int64 {
	return (time.Now().UnixNano() / int64(time.Millisecond)) - 1000
}

func IsValidTimeUnit(unit string) bool {
	units := []string{"1m", "3m", "5m", "15m", "30m", "1h", "2h", "4h", "6h", "8h", "12h", "1d", "3d", "1w", "1M"}
	return IsInStringList(unit, units)
}

func TimeLapseToIntervalCount(start int64, end int64, interval int64) (int64, error) {
	if end <= start {
		return 0, errors.New("start is higher than end")
	}
	timelapse := end - start
	return timelapse / interval, nil
}

func IntervalToMs(interval string) int64 {
	switch interval {
	case "1m":
		return 60000
	case "3m":
		return 180000
	case "5m":
		return 300000
	case "15m":
		return 900000
	case "30m":
		return 1800000
	case "1h":
		return 3600000
	case "2h":
		return 7200000
	case "4h":
		return 14400000
	case "6h":
		return 21600000
	case "8h":
		return 28800000
	case "12h":
		return 43200000
	case "1d":
		return 86400000
	case "3d":
		return 259200000
	case "1w":
		return 604800000
	case "1M":
		return 2419200000
	default:
		return 0
	}
}
