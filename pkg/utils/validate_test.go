package utils

import "testing"

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		email    string
		expected bool
		message  string
	}{
		{"example@example.com", true, "Email is valid"},
		{"user.name+tag+sorting@example.com", true, "Email is valid"},
		{"user_name@sub.domain.com", true, "Email is valid"},
		{"user-name@domain.co.uk", true, "Email is valid"},

		{"plainaddress", false, "Email must contain an '@' symbol"},
		{"@missingusername.com", false, "Email format is invalid"},
		{"username@.com", false, "Email format is invalid"},
		{"username@com", false, "Email must contain a '.' symbol"},
		{"username@domain.c", false, "Email format is invalid"},
		{"username@domain.com ", false, "Email must not contain spaces"},
		{" username@domain.com", false, "Email must not contain spaces"},
		{"username@ domain.com", false, "Email must not contain spaces"},
		{"user name@domain.com", false, "Email must not contain spaces"},
		{"username@domain.com.", false, "Email format is invalid"},
	}

	for _, test := range tests {
		result, message := ValidateEmail(test.email)
		if result != test.expected || message != test.message {
			t.Errorf("ValidateEmail(%q) = (%v, %q); want (%v, %q)", test.email, result, message, test.expected, test.message)
		}
	}
}

func TestValidateTime(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
		message  string
	}{
		{"2024-07-27T14:30:00Z", true, "Time is valid"},
		{"2024-07-27T14:30:00+01:00", true, "Time is valid"},
		{"2024-07-27T14:30:00-07:00", true, "Time is valid"},

		{"2024-07-27 14:30:00", false, "Time format is invalid"},
		{"27-07-2024T14:30:00Z", false, "Time format is invalid"},
		{"2024/07/27T14:30:00Z", false, "Time format is invalid"},
		{"InvalidTimeString", false, "Time format is invalid"},
	}

	for _, test := range tests {
		valid, msg, _ := ValidateTime(test.input)
		if valid != test.expected || msg != test.message {
			t.Errorf("ValidateTime(%q) = (%v, %q), expected (%v, %q)",
				test.input, valid, msg, test.expected, test.message)
		}
	}
}
