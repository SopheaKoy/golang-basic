package function

// errors is package
import (
	"errors"
	"fmt"
	"math/rand"
)

// Return and handle an error
func Information(name string) (string, error) {
	
	fmt.Println(randomFormat(), name)
	
	// way 1
	// if name == "" {
	// 	return "", errors.New("name should be not empty")
	// }
	// return name, nil

	// way 2
	return func() (string, error) {
		if name == "" {
			return "", errors.New("name must not be empty")
		}
		return name, nil
	}()
}

func randomFormat() string {
	formats := []string{
		"S",
		"O",
		"P",
		"H",
		"E",
		"A",
	}
	return formats[rand.Intn(len(formats))]
}
