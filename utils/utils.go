package utils

import (
	"unicode"
)

// CalculatePasswordStrength calculates the minimum number of steps required to make a password strong.
func CalculatePasswordStrength(password string) int {
	var steps int

	// Check the length of the password
	length := len(password)
	if length < 6 {
		steps += 6 - length
	} else if length > 20 {
		// Need to remove characters
		steps += length - 20
	}

	// Check for character types
	hasLower, hasUpper, hasDigit := false, false, false
	for _, ch := range password {
		switch {
		case unicode.IsLower(ch):
			hasLower = true
		case unicode.IsUpper(ch):
			hasUpper = true
		case unicode.IsDigit(ch):
			hasDigit = true
		}
	}

	// Count missing character types
	missingTypes := 0
	if !hasLower {
		missingTypes++
	}
	if !hasUpper {
		missingTypes++
	}
	if !hasDigit {
		missingTypes++
	}

	// Find the repeating character issues
	repeatSteps := 0
	repeatCount := 2 // To handle "aaa" scenario, we start at 2
	for i := 2; i < length; i++ {
		if password[i] == password[i-1] && password[i] == password[i-2] {
			repeatCount++
			repeatSteps = max(repeatSteps, repeatCount/3)
		} else {
			repeatCount = 2
		}
	}

	// Return the maximum of the necessary steps to handle length issues, missing types, and repeating characters
	return max(steps, max(missingTypes, repeatSteps))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
