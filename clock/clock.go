package clock

import "fmt"

//TestVersion is the exercism test version
const TestVersion = 2

//MinutesInADay is the constant number of minutes in 24 hours
const MinutesInADay = 1440

type clock struct {
	Hour   int
	Minute int
}

func Time(hourComponent, minuteComponent int) clock {
	timeObject := new(clock)
	hourComponent = (hourComponent + 24) % 24
	minuteComponent = (minuteComponent + MinutesInADay) % MinutesInADay
	return timeObject.recomputeTime((hourComponent * 60) + minuteComponent)
}

func (time *clock) String() string {
	return fmt.Sprintf("%02d:%02d", time.Hour, time.Minute)
}

func (time clock) recomputeTime(minutes int) clock {
	if minutes < 0 {
		minutes = MinutesInADay + minutes
	}
	if minutes >= 1440 {
		minutes = minutes - 1440
	}
	time.Hour = minutes / 60
	time.Minute = minutes % 60
	return time
}

//Add adds minutes(int) to a clock
func (time clock) Add(minutes int) clock {
	return time.recomputeTime(minutes + (time.Hour * 60) + time.Minute)
}