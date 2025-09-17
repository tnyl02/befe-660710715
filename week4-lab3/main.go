package main

import (
	"errors"
	"fmt"
)

type Student struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Year int `json:"year"`
	GPA float64 `json:"gpa"`
}

func (s *Student) IsHonor() bool{
	return s.GPA >= 3.50
}

func (s *Student) Validate() error{
	if s.Name == "" {
		return errors.New("name is required")
	}
	if s.Year < 1 || s.Year > 4 {
		return errors.New("year must be between 1-4")
	}
	if s.GPA < 0 || s.GPA > 4 {
		return errors.New("gpa must be between 0-4")
	}
	return nil
}

func main(){
	// var st Student = Student{ID:"1", Name:"Thanyalak", Email:"dechnakhornchai_t@silpakorn.edu", Year: 3, GPA: 3.20}
	student := []Student{
		{ID:"1", Name:"Thanyalak", Email:"dechnakhornchai_t@silpakorn.edu", Year: 3, GPA: 3.20} ,
		{ID:"2", Name:"AAA", Email:"AAA_A@silpakorn.edu", Year: 4, GPA: 3.00} ,
	}

	newStudent := Student{ID:"3", Name:"BBB", Email:"BBB_B@silpakorn.edu", Year: 2, GPA: 2.70}
	student = append(student, newStudent)

	for i, student:= range student {
		fmt.Printf("%d Honor %v\n",i, student.IsHonor())
		fmt.Printf("%d Validation %v\n",i, student.Validate())
	}

	
}