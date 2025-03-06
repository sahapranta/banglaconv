package banglaconv

import (
	"fmt"
	"math"
	"strings"
)

var digitMap = map[string]string{
	"1": "১", "2": "২", "3": "৩", "4": "৪",
	"5": "৫", "6": "৬", "7": "৭", "8": "৮",
	"9": "৯", "0": "০",
}

var numericWords = map[string]string{
	".":   "দশমিক",
	"0":   "শূন্য",
	"00":  "",
	"1":   "এক",
	"01":  "এক",
	"2":   "দুই",
	"02":  "দুই",
	"3":   "তিন",
	"03":  "তিন",
	"4":   "চার",
	"04":  "চার",
	"5":   "পাঁচ",
	"05":  "পাঁচ",
	"6":   "ছয়",
	"06":  "ছয়",
	"7":   "সাত",
	"07":  "সাত",
	"8":   "আট",
	"08":  "আট",
	"9":   "নয়",
	"09":  "নয়",
	"10":  "দশ",
	"11":  "এগারো",
	"12":  "বার",
	"13":  "তেরো",
	"14":  "চৌদ্দ",
	"15":  "পনের",
	"16":  "ষোল",
	"17":  "সতের",
	"18":  "আঠার",
	"19":  "উনিশ",
	"20":  "বিশ",
	"21":  "একুশ",
	"22":  "বাইশ",
	"23":  "তেইশ",
	"24":  "চব্বিশ",
	"25":  "পঁচিশ",
	"26":  "ছাব্বিশ",
	"27":  "সাতাশ",
	"28":  "আঠাশ",
	"29":  "ঊনত্রিশ",
	"30":  "ত্রিশ",
	"31":  "একত্রিশ",
	"32":  "বত্রিশ",
	"33":  "তেত্রিশ",
	"34":  "চৌত্রিশ",
	"35":  "পঁয়ত্রিশ",
	"36":  "ছত্রিশ",
	"37":  "সাঁইত্রিশ",
	"38":  "আটত্রিশ",
	"39":  "ঊনচল্লিশ",
	"40":  "চল্লিশ",
	"41":  "একচল্লিশ",
	"42":  "বিয়াল্লিশ",
	"43":  "তেতাল্লিশ",
	"44":  "চুয়াল্লিশ",
	"45":  "পঁয়তাল্লিশ",
	"46":  "ছেচল্লিশ",
	"47":  "সাতচল্লিশ",
	"48":  "আটচল্লিশ",
	"49":  "ঊনপঞ্চাশ",
	"50":  "পঞ্চাশ",
	"51":  "একান্ন",
	"52":  "বায়ান্ন",
	"53":  "তিপ্পান্ন",
	"54":  "চুয়ান্ন",
	"55":  "পঞ্চান্ন",
	"56":  "ছাপ্পান্ন",
	"57":  "সাতান্ন",
	"58":  "আটান্ন",
	"59":  "ঊনষাট",
	"60":  "ষাট",
	"61":  "একষট্টি",
	"62":  "বাষট্টি",
	"63":  "তেষট্টি",
	"64":  "চৌষট্টি",
	"65":  "পঁয়ষট্টি",
	"66":  "ছেষট্টি",
	"67":  "সাতষট্টি",
	"68":  "আটষট্টি",
	"69":  "ঊনসত্তর",
	"70":  "সত্তর",
	"71":  "একাত্তর",
	"72":  "বাহাত্তর",
	"73":  "তিয়াত্তর",
	"74":  "চুয়াত্তর",
	"75":  "পঁচাত্তর",
	"76":  "ছিয়াত্তর",
	"77":  "সাতাত্তর",
	"78":  "আটাত্তর",
	"79":  "ঊনআশি",
	"80":  "আশি",
	"81":  "একাশি",
	"82":  "বিরাশি",
	"83":  "তিরাশি",
	"84":  "চুরাশি",
	"85":  "পঁচাশি",
	"86":  "ছিয়াশি",
	"87":  "সাতাশি",
	"88":  "আটাশি",
	"89":  "ঊননব্বই",
	"90":  "নব্বই",
	"91":  "একানব্বই",
	"92":  "বিরানব্বই",
	"93":  "তিরানব্বই",
	"94":  "চুরানব্বই",
	"95":  "পঁচানব্বই",
	"96":  "ছিয়ানব্বই",
	"97":  "সাতানব্বই",
	"98":  "আটানব্বই",
	"99":  "নিরানব্বই",
	"100": "একশ",
}

// ToBengaliNumber converts English numerals to Bengali numerals
func ToBengaliNumber(numericText interface{}) string {
	if numericText == nil {
		return ""
	}

	str := fmt.Sprintf("%v", numericText)
	return replaceDigits(str)
}

/**
 * ToBengaliWord converts a number to its Bengali word representation.
 * It accepts various numeric types (int, float, string) as input.
 * Returns the Bengali word equivalent and an error if conversion fails.
 */
func ToBengaliWord(number interface{}) (string, error) {
	if number == nil || number == "" {
		return "", nil
	}

	num, err := convertToFloat64(number)

	if err != nil {
		return "", err
	}

	if num == 0 {
		return "শূন্য", nil
	}

	isFloat := num != math.Floor(num)

	var fractionPart string

	integerWords := ""

	if num < 0 {
		integerWords = "ঋণাত্মক "
		num = num * -1 // Make num positive
	}

	integerPart := int(num)
	integerWords = integerWords + integerToWords(integerPart)

	if isFloat {
		fractionStr := fmt.Sprintf("%.10f", num)
		parts := strings.Split(fractionStr, ".")
		if len(parts) == 2 {
			fractionPart = strings.TrimRight(parts[1], "0")
			if fractionPart == "" {
				fractionPart = "0"
			}
		}
		fractionWords := convertFractionToWords(fractionPart)
		return fmt.Sprintf("%s দশমিক %s", integerWords, fractionWords), nil
	}

	return integerWords, nil
}

func convertToFloat64(num interface{}) (float64, error) {
	switch v := num.(type) {
	case int:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case float32:
		return float64(v), nil
	case float64:
		return v, nil
	case string:
		var f float64
		_, err := fmt.Sscanf(v, "%f", &f)
		if err != nil {
			return 0, fmt.Errorf("invalid number format: %v", v)
		}
		return f, nil
	default:
		return 0, fmt.Errorf("unsupported type: %T", num)
	}
}

func replaceDigits(str string) string {
	result := ""
	str = strings.TrimLeft(str, "0")
	for _, char := range str {
		if replacement, exists := digitMap[string(char)]; exists {
			result += replacement
		} else {
			result += string(char)
		}
	}
	return result
}

func convertFractionToWords(fraction string) string {
	frac := strings.TrimLeft(fraction, "0")

	if frac == "" {
		return ""
	}

	var words []string

	for _, digit := range fraction {
		words = append(words, numericWords[string(digit)])
	}

	return strings.TrimSpace(strings.Join(words, " "))
}

func integerToWords(num int) string {
	if num == 0 {
		return "শূন্য"
	}

	words := []string{}

	if num >= 10000000 {
		crore := num / 10000000
		if crore > 0 {
			words = append(words, fmt.Sprintf("%s কোটি", integerToWords(crore)))
		}
		num %= 10000000
	}

	if num >= 100000 {
		lac := num / 100000
		if lac > 0 {
			words = append(words, fmt.Sprintf("%s লক্ষ", integerToWords(lac)))
		}
		num %= 100000
	}

	if num >= 1000 {
		thousand := num / 1000
		if thousand > 0 {
			words = append(words, fmt.Sprintf("%s হাজার", integerToWords(thousand)))
		}
		num %= 1000
	}

	if num >= 100 {
		hundred := num / 100
		if hundred > 0 {
			words = append(words, fmt.Sprintf("%sশ", integerToWords(hundred)))
		}
		num %= 100
	}

	if num > 0 {
		words = append(words, parseRemainingNumber(num))
	}

	return strings.Join(words, " ")
}

func parseRemainingNumber(num int) string {
	if num == 0 {
		return ""
	}

	if num <= 99 {
		return numericWords[fmt.Sprintf("%d", num)]
	}

	tensDigit := (num / 10) * 10
	onesDigit := num % 10

	if onesDigit == 0 {
		return numericWords[fmt.Sprintf("%d", tensDigit)]
	}

	tensWord := numericWords[fmt.Sprintf("%d", tensDigit)]
	onesWord := numericWords[fmt.Sprintf("%d", onesDigit)]

	if tensWord == "" {
		return onesWord
	}
	if onesWord == "" {
		return tensWord
	}

	return fmt.Sprintf("%s %s", tensWord, onesWord)
}
