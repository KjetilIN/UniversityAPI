package serverstats

import (
    "time"
)

//start time of when service was started 
var startTime time.Time

//Get the time since the service started
func GetUptime() time.Duration {
    return time.Since(startTime)
}

//Set the start time to current time. 
func InitServerTimer() {
    startTime = time.Now()
}


