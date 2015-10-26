package clock

import "fmt"

//TestVersion is the exercism test version.
const TestVersion = 2

//MinutesInADay is the constant number of minutes in 24 hours.
const MinutesInADay = 1440

//Clock holds the time as hour & minute.
type Clock struct {
	//un-exported fields inside exported struct. Problem!?
	hour   int
	minute int
}

//Time is a 'constructor' that creates and initialises a Clock.
func Time(hourComponent, minuteComponent int) Clock {
	clockObject := new(Clock)

	// trying to clean up input for negatives & overflows(beyond 24 & 1440)
	hourComponent = (hourComponent + 24) % 24
	minuteComponent = (minuteComponent + MinutesInADay) % MinutesInADay
	return clockObject.setTimeInClock((hourComponent * 60) + minuteComponent)
}

func (time *Clock) String() string {
	return fmt.Sprintf("%02d:%02d", time.hour, time.minute)
}

//function on Clock to compute hour and minute from a total minutes(since 00:00).
func (time Clock) setTimeInClock(minutes int) Clock {
	if minutes < 0 {
		minutes = MinutesInADay + minutes
	}
	if minutes >= 1440 {
		minutes = minutes - 1440
	}
	time.hour = minutes / 60
	time.minute = minutes % 60
	return time
}

//Add adds minutes to the time on a Clock.
func (time Clock) Add(minutes int) Clock {
	return time.setTimeInClock(minutes + (time.hour * 60) + time.minute)
}
