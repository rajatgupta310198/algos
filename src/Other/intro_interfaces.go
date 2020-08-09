package main

import "fmt"

type MarkSheetProcessing interface {
	Add(subMark int) int;
	Percentage(totalMarks, numberOfSubjects int) float64;
}
type Student struct {
	name string;
}

func (s *Student) Add(subMark int) int  {
	return 1 + subMark;
}

func (s *Student) Percentage(totalMarks, numberOfSubjects int) float64  {
	return float64(totalMarks / numberOfSubjects)
}



func main() {
	var m MarkSheetProcessing;
	m = &Student{name:"Rajat"}
	fmt.Print(m.Add(1))

}
