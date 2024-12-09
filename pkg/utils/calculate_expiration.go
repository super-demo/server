package utils

import "time"

func CalculateExpiration(duration int) int64 {
	return time.Now().Add(time.Minute * time.Duration(duration)).Unix()
}
