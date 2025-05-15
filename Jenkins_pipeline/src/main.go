package main

import (
	"fmt"
	"os"
	"regexp"
)

// ValidatePhoneNumber ensures exactly 10 digits, no letters, no symbols
func ValidatePhoneNumber(phone string) bool {
	re := regexp.MustCompile(`^[0-9]{10}$`)
	return re.MatchString(phone)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./main <phone-number>")
		return
	}

	phone := os.Args[1]

	if ValidatePhoneNumber(phone) {
		fmt.Println("Valid phone number")
	} else {
		fmt.Println("Invalid phone number")
	}
}
