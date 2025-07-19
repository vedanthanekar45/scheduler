package main

import (
	"fmt"
	"time"
)

type Teacher struct {
	Name string
	Sub Subject 
	Start_shift time.Time
	End_shift time.Time
}

type Subject struct {
	Name string
	hours int
}

type Timeslot struct {
	start time.Time
	end time.Time
}

type Day struct {
	periods []Timeslot
}

type Timetable struct {
	week []Day
}