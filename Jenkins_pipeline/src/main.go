package main

import (
	"regexp"
	"fmt"
)

// ValidatePhoneNumber checks if the provided phone number is valid based on a simple pattern
func ValidatePhoneNumber(phone string) bool {
	// A simple regex for phone numbers: starts with a + and followed by 10-15 digits
	// Adjust the regex as per your specific phone number format
	re := regexp.MustCompile(`^\+?[0-9]{10,15}$`)
	return re.MatchString(phone)
}

func main() {
	// Example usage:
	phone := "+12345678901"
	if ValidatePhoneNumber(phone) {
		fmt.Println("Valid phone number")
	} else {
		fmt.Println("Invalid phone number")
	}
}
