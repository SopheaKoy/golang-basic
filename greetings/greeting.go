package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// return without handle error
func Hello(name string) string {
	// Declare variables
	var gender string
	var age int
	var isOpen bool
	var message string

	// Assign values for demonstration
	gender = "Male"
	age    = 25
	isOpen = true

	// Example 1: Basic formatting with variables
	message = fmt.Sprintf("Hi, %v! Gender: %s, Age: %d, Is Open: %t", name, gender, age, isOpen)

	// Example 2: Reordering arguments using %[n]d
	// This line will overwrite the previous message
	// message = fmt.Sprintf("%[2]d %[1]d", 11, 22)

	// Shortcut version (uncomment to use):
	// message := fmt.Sprintf("Hi, %v welcome!", name)

	return message
}

func Hi(name string) (string, error) {

	msg := fmt.Sprintf("Hi, My name is %v", name)
	return msg, nil
}

func Greeting(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return name, errors.New("empty name")
    }
    // Create a message using a random format.
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
    // A slice of message formats.
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    // Return a randomly selected message format by specifying
    // a random index for the slice of formats.
    return formats[rand.Intn(len(formats))]
}