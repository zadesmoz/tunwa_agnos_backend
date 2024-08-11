package utils

import "testing"

func TestCalculatePasswordStrength(t *testing.T) {

	tests := []struct {
		name     string
		password string
		expected int
	}{
		{
			name:     "Short password, missing character types",
			password: "aA1",
			expected: 3,
		},
		{
			name:     "Password meets all criteria",
			password: "1445D1cd",
			expected: 0,
		},
		{
			name:     "Password is too short and has repeating characters",
			password: "aaaaa",
			expected: 2,
		},
		{
			name:     "Minimum length password with all character types",
			password: "A1b!@#",
			expected: 0,
		},
		{
			name:     "Password with 3 repeating characters in a row",
			password: "aaaAAA1",
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CalculatePasswordStrength(tt.password)
			if result != tt.expected {
				t.Errorf("CalculatePasswordStrength(%q) = %d; want %d", tt.password, result, tt.expected)
			}
		})
	}
}
