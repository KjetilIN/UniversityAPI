package serverstats

import (
	"testing"
	"time"
)

func TestInitTimerWorks(t *testing.T) {

	//Init server
	InitServerTimer();

	//We wait two seconds
	time.Sleep(time.Second*2);

	//Check if the start time is within the range 
	if(startTime.Second() > 2 && startTime.Second() < 5){
		t.Fatal("Server did not start correctly")
		t.Error()
	}

}


func TestGetUptime(t *testing.T) {
    // Set the start time to a known value (5 minutes ago)
    startTime = time.Now().Add(-5 * time.Minute)

    // Call GetUptime and check that the returned duration is within 1 second of the expected value (5 minutes)
    expectedDuration := 5 * time.Minute
    actualDuration := GetUptime()
    if actualDuration < expectedDuration-time.Second || actualDuration > expectedDuration+time.Second {
        t.Errorf("GetUptime() = %v, expected %v", actualDuration, expectedDuration)
    }
}
