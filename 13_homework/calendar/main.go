package main

import (
	"time"
)

type Calendar struct {
	time.Time
}

func (c *Calendar) CurrentQuarter() int {
	month := int(c.Time.Month())
	for month%3 != 0 {
		month++
	}
	return month / 3
}

func NewCalendar(time time.Time) Calendar {
	return Calendar{time}
}

func main() {
}
