package serverstats

import (
    "time"
)

//start time of when service was started 
var startTime time.Time

//Get the time sice the service started
func Uptime() time.Duration {
    return time.Since(startTime)
}

//Set the start time to current time. 
func Init() {
    startTime = time.Now()
}


