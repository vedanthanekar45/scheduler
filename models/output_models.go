package models

import (
	"time"
)

type OutputTeacher struct {
	Name       string
	Sub        OutputSubject
	StartShift time.Time
	EndShift   time.Time
}

type OutputSubject struct {
	Name  string
	Hours int
}

type OutputTimeslot struct {
	Start time.Time
	End   time.Time
}

type Day struct {
	Periods []OutputTimeslot
}

type OutputTimetable struct {
	Week []Day
}
