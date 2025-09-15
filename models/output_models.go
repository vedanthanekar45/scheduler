package models

type ScheduledPeriod struct {
	ClassGroupName string
	Subject        InputSubject
	Teacher        InputTeacher
	Timeslot       InputTimeslot
}

type TimetableSolution []ScheduledPeriod
