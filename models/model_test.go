package models

import (
	"fmt"
	"testing"
	"time"
)

func TestTeacherInit(t *testing.T) {
	subject := OutputSubject{Name: "Math", Hours: 5}
	start := time.Date(2025, 7, 18, 9, 0, 0, 0, time.Local)
	end := time.Date(2025, 7, 18, 13, 0, 0, 0, time.Local)

	teacher := OutputTeacher{
		Name:       "Cassian Andor",
		Sub:        subject,
		StartShift: start,
		EndShift:   end,
	}

	fmt.Printf("Teacher: %s, Subject: %s, Start: %s, End: %s\n",
		teacher.Name, teacher.Sub.Name, teacher.StartShift.Format(time.RFC3339),
		teacher.EndShift.Format(time.RFC3339))

	if teacher.Sub.Hours != 5 {
		t.Errorf("Expected 5 hours, got %d", teacher.Sub.Hours)
	}

	if teacher.StartShift.After(teacher.EndShift) {
		t.Error("Start time is after end time!")
	}

}
