package models

type ProblemDefinition struct {
	ClassGroupName string          `json:"classGroupName"`
	Subjects       []InputSubject  `json:"subjects"`
	Teachers       []InputTeacher  `json:"teachers"`
	Timeslots      []InputTimeslot `json:"timeslots"`
}

type InputSubject struct {
	Name           string `json:"name"`
	PeriodsPerWeek int    `json:"periodsPerWeek"`
}

type InputTeacher struct {
	Name              string   `json:"name"`
	TeachableSubjects []string `json:"teachableSubjects"`
}

type InputTimeslot struct {
	Day    string `json:"day"`
	Period int    `json:"period"`
}
