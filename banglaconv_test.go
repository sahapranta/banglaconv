package banglaconv

import (
	"fmt"
	"testing"
)

func TestToBengaliNumber(t *testing.T) {
	converter := &NumConverter{}
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
			result := converter.ToBengaliNumber(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}

func TestToBengaliWord(t *testing.T) {
	converter := &NumConverter{}
	testCases := []struct {
		input    interface{}
		expected string
	}{
		{0, "শুন্য"},
		{1, "এক"},
		{10, "দশ"},
		{34, "চৌত্রিশ"},
		{100, "একশত"},
		{1234, "এক হাজার দুইশত চৌত্রিশ"},
		{1234567, "বার লক্ষ চৌত্রিশ হাজার পাঁচশত সাতষট্টি"},
		{1234.56, "এক হাজার দুইশত চৌত্রিশ দশমিক পাঁচ ছয়"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Convert %v", tc.input), func(t *testing.T) {
			result, err := converter.ToBengaliWord(tc.input)
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
	converter := &NumConverter{}

	_, err := converter.ToBengaliWord("invalid")
	if err == nil {
		t.Errorf("Expected error for invalid input, got nil")
	}
}

func BenchmarkToBengaliNumber(b *testing.B) {
	converter := &NumConverter{}
	for i := 0; i < b.N; i++ {
		converter.ToBengaliNumber(1234567)
	}
}

func BenchmarkToBengaliWord(b *testing.B) {
	converter := &NumConverter{}
	for i := 0; i < b.N; i++ {
		converter.ToBengaliWord(1234567)
	}
}
