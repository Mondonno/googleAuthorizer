package gauth

import (
	"math"
)

func GetCurrentEpoch() (int, int) { // returns an epoch and count down for user
	epoch := (math.Round(getCurrentTime() / 1000.0))
	countDown := 30 - (epoch % 30)

	return epoch, countDown
}

func GetCurrentTime() int {
	return time.Now().UTC().UnixNano() / 1e6
}