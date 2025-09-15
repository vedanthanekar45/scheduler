package utils

import (
	"fmt"
	"os"
	"scheduler/v2/models"
)

func CheckValues(problemDef models.ProblemDefinition) {
	if problemDef.ClassGroupName == "" || len(problemDef.Subjects) == 0 || len(problemDef.Teachers) == 0 || len(problemDef.Timeslots) == 0 {
		fmt.Println("Error: Input JSON is missing one or more mandatory top-level fields (classGroupName, subjects, teachers, timeslots).")
		os.Exit(1)
	}
}
