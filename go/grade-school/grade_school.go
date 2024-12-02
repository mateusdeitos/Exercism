package school

import (
	"slices"
	"strings"
)

// Define the Grade and School types here.

type School struct {
	gradeMap map[int][]string
	grades   []int
}

type Grade struct {
	grade    int
	students []string
}

func New() *School {
	return &School{
		gradeMap: map[int][]string{},
		grades:   []int{},
	}
}

func (s *School) Add(student string, grade int) {
	if _, ok := s.gradeMap[grade]; !ok {
		s.gradeMap[grade] = []string{}
		s.grades = append(s.grades, grade)
	}

	s.gradeMap[grade] = append(s.gradeMap[grade], student)
}

func (s *School) Grade(level int) []string {
	return s.gradeMap[level]
}

func (s *School) Enrollment() []Grade {
	result := make([]Grade, len(s.grades))
	slices.Sort(s.grades)
	for i := 0; i < len(s.grades); i++ {
		grade := s.grades[i]
		students := s.gradeMap[grade]
		slices.Sort(students)

		result[i] = Grade{grade, []string{strings.Join(students, " ")}}
	}

	return result
}
