# Simple Task Scheduler Engine

A simple, constraint-based scheduler written in Go. This project uses a backtracking algorithm to solve Constraint Satisfaction Problems (CSPs), starting with a simplified school timetable model.

---

## What it Does

This program takes a problem definition—such as subjects, teachers, and time slots—from a JSON file and generates a valid, conflict-free schedule. The core logic is designed to be adaptable for other scheduling problems.

---

## Getting Started

### Prerequisites

* Go (version 1.18 or higher recommended)

### How to Run

1.  **Clone the repository:**
    ```sh
    git clone <your-repository-url>
    cd <your-repository-directory>
    ```

2.  **Create an input file:**
    Create a file named `input.json` to define your scheduling problem. See the example below.

3.  **Run the program:**
    Execute the scheduler from your terminal, providing the path to your input file.
    ```sh
    go run . --input=input.json
    ```

---

## Example Usage

This example shows how to generate a timetable for a single class, "Grade 9A".

### Input (`input.json`)

This file defines the subjects to be scheduled, the teachers available to teach them, and the available time slots.

```json
{
  "classGroupName": "Grade 9A",
  "subjects": [
    { "name": "Mathematics", "periodsPerWeek": 4 },
    { "name": "Science", "periodsPerWeek": 4 },
    { "name": "History", "periodsPerWeek": 3 }
  ],
  "teachers": [
    {
      "name": "Mr. Sharma",
      "teachableSubjects": ["Mathematics"]
    },
    {
      "name": "Ms. Gupta",
      "teachableSubjects": ["Science"]
    },
    {
      "name": "Mr. Singh",
      "teachableSubjects": ["History", "Science"]
    }
  ],
  "timeslots": [
    { "day": "Monday", "period": 1 },
    { "day": "Monday", "period": 2 },
    { "day": "Monday", "period": 3 },
    { "day": "Tuesday", "period": 1 }
  ]
}