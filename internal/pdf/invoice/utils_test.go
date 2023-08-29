package pdf_invoice

import (
	"testing"
)

func TestFormatNumber(t *testing.T) {
	tests := []struct {
		input    float64
		expected string
	}{
		{23370.00, "23 370,00"},
		{1000000.00, "1 000 000,00"},
		{1234.56, "1 234,56"},
		{7.50, "7,50"},
		{9876543210.01, "9 876 543 210,01"},
	}

	for _, test := range tests {
		got := formatSumToString(test.input)
		if got != test.expected {
			t.Errorf("Expected '%s' but got '%s' for input %f", test.expected, got, test.input)
		}
	}
}

func TestNumberToPolish(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{0, "zero"},
		{5, "pięć"},
		{11, "jedenaście"},
		{123, "sto dwadzieścia trzy"},
		{1001, "jeden tysiąc jeden"},
		{12345, "dwanaście tysięcy trzysta czterdzieści pięć"},
		{20000, "dwadzieścia tysięcy"},
		{23370, "dwadzieścia trzy tysiące trzysta siedemdziesiąt"},
		{99999, "dziewięćdziesiąt dziewięć tysięcy dziewięćset dziewięćdziesiąt dziewięć"},
	}

	for _, test := range tests {
		got := numberToPolish(test.input)
		if got != test.expected {
			t.Errorf("Expected '%s' but got '%s' for input %d", test.expected, got, test.input)
		}
	}
}

func TestSplitFloat(t *testing.T) {
	tests := []struct {
		input     float64
		wantInt   int
		wantFloat string
	}{
		{123.456, 123, "46"},
		{-123.456, -123, "46"},
		{0, 0, "00"},
		{123, 123, "00"},
		{0.99, 0, "99"},
		{0.9999, 1, "00"},
		{23370.00, 23370, "00"},
	}

	for _, tt := range tests {
		gotInt, gotFloat := splitFloat(tt.input)
		if gotInt != tt.wantInt || gotFloat != tt.wantFloat {
			t.Errorf("splitFloat(%f) = %d, %s; want %d, %s", tt.input, gotInt, gotFloat, tt.wantInt, tt.wantFloat)
		}
	}
}

func TestGetLastDayOfMonth(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		hasError bool
	}{
		{"31-08-2023", "31-08-2023", false},
		{"15-02-2020", "29-02-2020", false}, // Leap year
		{"15-02-2021", "28-02-2021", false}, // Non-leap year
		{"28-02-2021", "28-02-2021", false},
		{"01-01-2023", "31-01-2023", false},
		{"01-04-2023", "30-04-2023", false},
		{"32-01-2023", "", true}, // Invalid date
		{"15-13-2023", "", true}, // Invalid month
	}

	for _, test := range tests {
		result, err := getLastDayOfMonth(test.input)
		if test.hasError && err == nil {
			t.Errorf("Expected error for input %v, but got none", test.input)
			continue
		}
		if !test.hasError && err != nil {
			t.Errorf("Didn't expect error for input %v, but got: %v", test.input, err)
			continue
		}
		if result != test.expected {
			t.Errorf("For input %v, expected %v but got %v", test.input, test.expected, result)
		}
	}
}
