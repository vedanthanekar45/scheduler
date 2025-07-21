package models

import (
	"time"
)

type Teacher struct {
	Name       string
	Sub        Subject
	StartShift time.Time
	EndShift   time.Time
}

type Subject struct {
	Name  string
	Hours int
}

type Timeslot struct {
	Start time.Time
	End   time.Time
}

type Day struct {
	Periods []Timeslot
}

type Timetable struct {
	Week []Day
}
