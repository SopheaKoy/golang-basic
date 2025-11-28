package variable

import "fmt"

/*
	Noted: If variable need to use outside. Pls Declaration with Capital Letter
*/

// constants
const PI 	   = 3.14;
const Greeting = "Bong Phea";

// declare variable
var Name = "Sophea";
var Age	 = 2025;

// 4. Declare multiple variables
var X, Y, Z int = 1, 2, 3
var FirstName, Grade, Active = "Phea", 25, true

// 5. Grouped variable declaration
var (
    Username string = "admin"
    Password string = "secret"
    Port     int    = 8080
)

// 6. Zero values
var Number   int
var Refernce string
var IsActive bool
var Balance  float64

var Firstname 	string
var Lastname 	string
var Gender		string
var Statuscode	int
var Isupdated	bool
var Filename	string

func Show() {
	fmt.Println("Hello kon papa.")
}