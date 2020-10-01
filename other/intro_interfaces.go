package main

import "fmt"

// MarkSheetProcessing Interface
type MarkSheetProcessing interface {
	Add(subMark int) int
	Percentage(totalMarks, numberOfSubjects int) float64
}

// Student struct
type Student struct {
	name string
}

// Add to add mark
func (s *Student) Add(subMark int) int {
	return 1 + subMark
}

// Percentage calculation
func (s *Student) Percentage(totalMarks, numberOfSubjects int) float64 {
	return float64(totalMarks / numberOfSubjects)
}

func main() {
	var m MarkSheetProcessing
	m = &Student{name: "Rajat"}
	fmt.Print(m.Add(1))

}
