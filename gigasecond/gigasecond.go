package gigasecond
import (
	"time"
	"math"
)

//TestVersion is the exercism test version
const TestVersion = 2

var Birthday time.Time

func AddGigasecond(timeObject time.Time) time.Time {
	timeObject = timeObject.Add(time.Duration(int(math.Pow(10, 9))) * time.Second)
	return timeObject
}
