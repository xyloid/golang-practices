package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
	var gs time.Duration = 1000000000000000000
	return t.Add(gs)
}
