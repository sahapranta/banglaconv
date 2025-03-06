package banglaconv

import (
	"fmt"
	"math"
	"testing"
)

func TestToBengaliNumber(t *testing.T) {
	testCases := []struct {
		input    interface{}
		expected string
	}{
		{1234, "১২৩৪"},
		{0, "০"},
		{"5678", "৫৬৭৮"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Convert %v", tc.input), func(t *testing.T) {
			result := ToBengaliNumber(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}

func TestToBengaliWord(t *testing.T) {
	testCases := []struct {
		input    interface{}
		expected string
	}{
		{1, "এক"},
		{10, "দশ"},
		{34, "চৌত্রিশ"},
		{0, "শূন্য"},
		{100, "একশ"},
		{1234, "এক হাজার দুইশ চৌত্রিশ"},
		{1234567, "বার লক্ষ চৌত্রিশ হাজার পাঁচশ সাতষট্টি"},
		{1234.56, "এক হাজার দুইশ চৌত্রিশ দশমিক পাঁচ ছয়"},
		{0.05, "শূন্য দশমিক শূন্য পাঁচ"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Convert %v", tc.input), func(t *testing.T) {
			result, err := ToBengaliWord(tc.input)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}

func TestInvalidInput(t *testing.T) {
	_, err := ToBengaliWord("invalid")
	if err == nil {
		t.Errorf("Expected error for invalid input, got nil")
	}
}

func TestToBengaliNumberEdgeCases(t *testing.T) {
	testCases := []struct {
		input    interface{}
		expected string
	}{
		{-1234, "-১২৩৪"},
		{uint(789), "৭৮৯"},
		{"000123", "১২৩"},
		{"", ""},
		{nil, ""},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Convert %v", tc.input), func(t *testing.T) {
			result := ToBengaliNumber(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}

func TestConvertFractionToWords(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"-123", "এক দুই তিন"},
		{"123", "এক দুই তিন"},
		{"000", ""},
		{"001", "শূন্য শূন্য এক"},
		{"0.0001", "শূন্য দশমিক শূন্য শূন্য শূন্য এক"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Convert fraction %s", tc.input), func(t *testing.T) {
			result := convertFractionToWords(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}

func TestConvertToFloat64(t *testing.T) {
	testCases := []struct {
		input    interface{}
		expected float64
		hasError bool
	}{
		{int(123), 123.0, false},
		{int32(456), 456.0, false},
		{int64(789), 789.0, false},
		{float32(1.23), 1.23, false},
		{float64(4.56), 4.56, false},
		{"7.89", 7.89, false},
		{"invalid", 0, true},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Convert %v", tc.input), func(t *testing.T) {
			result, err := convertToFloat64(tc.input)
			if tc.hasError {
				if err == nil {
					t.Errorf("Expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				// Use math.Abs to compare floating-point numbers with small tolerance
				if math.Abs(result-tc.expected) > 1e-6 {
					t.Errorf("Expected %f, got %f", tc.expected, result)
				}
			}
		})
	}
}

func TestToBengaliWordEdgeCases(t *testing.T) {
	testCases := []struct {
		input    interface{}
		expected string
	}{
		{-1, "ঋণাত্মক এক"},
		{-123, "ঋণাত্মক একশ তেইশ"},
		{0.0001, "শূন্য দশমিক শূন্য শূন্য শূন্য এক"},
		{0.000, "শূন্য"},
		{123.00, "একশ তেইশ"},
		{12300000, "এক কোটি তেইশ লক্ষ"},
		{12300000000, "এক হাজার দুইশ ত্রিশ কোটি"},
		{"000123", "একশ তেইশ"},
		{"", ""},
		{nil, ""},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Convert %v", tc.input), func(t *testing.T) {
			result, err := ToBengaliWord(tc.input)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}

func TestParseRemainingNumber(t *testing.T) {
	testCases := []struct {
		input    int
		expected string
	}{
		{0, ""},
		{15, "পনের"},
		{25, "পঁচিশ"},
		{34, "চৌত্রিশ"},
		{45, "পঁয়তাল্লিশ"},
		{451, "এক"},
		{450, ""},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Parse %d", tc.input), func(t *testing.T) {
			result := parseRemainingNumber(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}

func TestIntegerToWords(t *testing.T) {
	testCases := []struct {
		input    int
		expected string
	}{
		{0, "শূন্য"},
		{45, "পঁয়তাল্লিশ"},
		{100, "একশ"},
		{1234, "এক হাজার দুইশ চৌত্রিশ"},
		{1000000, "দশ লক্ষ"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Convert %d", tc.input), func(t *testing.T) {
			result := integerToWords(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}

func TestInvalidInputCases(t *testing.T) {
	testCases := []interface{}{
		true, false, struct{}{}, []int{1, 2, 3},
	}

	for _, input := range testCases {
		t.Run(fmt.Sprintf("Invalid %v", input), func(t *testing.T) {
			_, err := ToBengaliWord(input)
			if err == nil {
				t.Errorf("Expected error for input %v, got nil", input)
			}
		})
	}
}

func TestConvertToFloat64EdgeCases(t *testing.T) {
	testCases := []struct {
		input    interface{}
		hasError bool
	}{
		{uint(123), false},
		{nil, true},
		{true, true},
		{false, true},
		{complex(1, 2), true},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Convert %v", tc.input), func(t *testing.T) {
			_, err := convertToFloat64(tc.input)
			if tc.hasError && err == nil {
				t.Errorf("Expected error, got nil")
			}
		})
	}
}

func BenchmarkToBengaliNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToBengaliNumber(1234567)
	}
}

func BenchmarkToBengaliWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToBengaliWord(1234567)
	}
}
