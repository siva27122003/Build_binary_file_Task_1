package main

import (
	"fmt"
	"os"
	"regexp"
)

// ValidatePhoneNumber checks if the provided phone number is valid based on a simple pattern
func ValidatePhoneNumber(phone string) bool {
	// Regex: optional +, followed by 10 to 15 digits
	re := regexp.MustCompile(`^\+?[0-9]{10,15}$`)
	return re.MatchString(phone)
}

func main() {
	// Check if a phone number was passed as an argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./main <phone-number>")
		return
	}

	// Get the phone number from the command-line arguments
	phone := os.Args[1]

	if ValidatePhoneNumber(phone) {
		fmt.Println("Valid phone number")
	} else {
		fmt.Println("Invalid phone number")
	}
}
