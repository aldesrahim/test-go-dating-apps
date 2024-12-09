package utils

import "time"

func GetAgeFromDate(dt time.Time) int {
	return int(time.Now().Sub(dt).Hours() / (24 * 365)) // calculate age in years
}
