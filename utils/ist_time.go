package utils

import "time"

var ISTLocation, _ = time.LoadLocation("Asia/Kolkata")

func GetCurrentTimeIST() time.Time {
	return time.Now().In(ISTLocation)
}
