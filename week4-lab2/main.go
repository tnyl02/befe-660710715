package main

import (
	"fmt"
)
// var email string = "dechnakhornchai_t@silpakorn.edu"

func main()  {
	// var name string = "Thanyalak"
	var age int = 21
	email := "dechnakhornchai_t@silpakorn.edu"
	gpa := 3.20
	firstname, lastname := "Thanyalak" , "Dechnakhornchai"

	fmt.Printf("Name %s %s, age %d, email %s, gpa %.2f\n", firstname, lastname, age, email, gpa)
}