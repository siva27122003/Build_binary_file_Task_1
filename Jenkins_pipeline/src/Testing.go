package main

import (
	"testing"
)

// TestValidatePhoneNumber tests the ValidatePhoneNumber function
func TestValidatePhoneNumber(t *testing.T) {
	tests := []struct {
		phone   string
		isValid bool
	}{
		{"+12345678901", true},       // valid phone number
		{"1234567890", true},         // valid phone number without '+'
		{"+12345", false},            // invalid phone number (too short)
		{"12345678901234567890", false}, // invalid phone number (too long)
		{"123-456-7890", false},      // invalid phone number (wrong format)
		{"abcdef12345", false},      // invalid phone number (letters included)
	}

	for _, tt := range tests {
		t.Run(tt.phone, func(t *testing.T) {
			if got := ValidatePhoneNumber(tt.phone); got != tt.isValid {
				t.Errorf("ValidatePhoneNumber(%s) = %v, want %v", tt.phone, got, tt.isValid)
			}
		})
	}
}
