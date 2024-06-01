package timeunit

type TimeUnit int

const (
	NANOS TimeUnit = iota
	MICROS
	MILLIS
	SECONDS
	MINUTES
	HOURS
	HALF_DAYS
	DAYS
)
