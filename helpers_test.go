package main

import (
	"testing"
)

func TestIsValidUrl(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"https://example.com", true},
		{"http://sub.domain.co.uk/path?query=123#fragment", true},
		{"ftp://user:pass@example.com:21", true},
		{"invalid-url", false},
		{"https://", false},
		{"", false},
		{"http://.com", false},
		{"http:// invalid .com", false},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := isValidUrl(test.input)
			if result != test.expected {
				t.Errorf("isValidUrl(%q) = %v; want %v", test.input, result, test.expected)
			}
		})
	}
}

func TestGenerateShortCode(t *testing.T) {
	tests := []struct {
		input string
	}{
		{"https://example.com"},
		{"http://test.com/path"},
		{""},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := generateShortCode(test.input)

			if len(result) != CodeSize {
				t.Errorf("generateShortCode(%q) length = %d; want %d", test.input, len(result), CodeSize)
			}

			code1 := generateShortCode(test.input)
			code2 := generateShortCode(test.input)
			if code1 != code2 {
				t.Errorf("generateShortCode(%q) not deterministic: %q != %q", test.input, code1, code2)
			}
		})
	}
}

func TestGenerateShortCodeDeterministic(t *testing.T) {
	input1 := "https://example.com"
	input2 := "https://test.com"
	if generateShortCode(input1) == generateShortCode(input2) {
		t.Errorf("generateShortCode produced same code for different inputs: %q and %q", input1, input2)
	}
}
