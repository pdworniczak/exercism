package gigasecond

import "time"

const GIGASECOND = 1000000000

func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * GIGASECOND)
}
