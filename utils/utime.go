package utils

import (
	"fmt"
	"time"
)

func CurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func CurrentDate() string {
	return time.Now().Format("20060102")
}

func CurrentDateByPlace(place string) string {
	return time.Now().Format(fmt.Sprintf("2006%s01%s02", place, place))
}

func FormatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
