package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Тест функции реверса строки
func TestReverseStrings(t *testing.T) {
	TBReverseStrings(t, nil)
}

func BenchmarkReverseStrings(b *testing.B) {
	TBReverseStrings(nil, b)
}

func TBReverseStrings(t *testing.T, b *testing.B) {
	testTable := []struct {
		name     string
		text     string
		expected string
	}{
		{
			name:     "reverse 2 short words",
			text:     "hello world",
			expected: "dlrow olleh",
		},
		{
			name:     "empty string",
			text:     "",
			expected: "",
		},
		{
			name:     "reverse medium string with 6 words",
			text:     "hello world this is a test",
			expected: "tset a si siht dlrow olleh",
		},
		{
			name:     "reverse medium string with 8 words",
			text:     "Lorem ipsum dolor sit amet consectetur adipiscing elit",
			expected: "tile gnicsipida rutetcesnoc tema tis rolod muspi meroL",
		},
	}

	if t != nil {
		for _, testCase := range testTable {

			result := ReverseStrings(testCase.text)

			t.Logf("Calling ReverseStrings(%s), result: %s", testCase.text, result)
			assert.Equal(t, testCase.expected, result,
				fmt.Sprintf("ReverseStrings returned %s, expected %s",
					result, testCase.expected))
		}
	}
	if b != nil {
		for _, testCase := range testTable {
			b.Run(testCase.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					ReverseStrings(testCase.text)
				}
			})
		}
	}
}

// Тест алгоритма сортировки слиянием
func TestMergeSort(t *testing.T) {
	TBMergeSort(t, nil)
}

func BenchmarkMergeSort(b *testing.B) {
	TBMergeSort(nil, b)
}

func TBMergeSort(t *testing.T, b *testing.B) {
	testTable := []struct {
		name           string
		firstArray     []int
		lenFirstArray  int
		secondArray    []int
		lenSecondArray int
		expected       []int
	}{
		{
			name:           "merge 2 arrays with 3 and 3 elements",
			firstArray:     []int{1, 2, 3, 0, 0, 0},
			lenFirstArray:  3,
			secondArray:    []int{2, 5, 6},
			lenSecondArray: 3,
			expected:       []int{1, 2, 2, 3, 5, 6},
		},
		{
			name:           "merge 2 arrays with 7 and 5 elements",
			firstArray:     []int{1, 5, 7, 11, 15, 21, 25, 0, 0, 0, 0, 0},
			lenFirstArray:  7,
			secondArray:    []int{6, 14, 29, 37, 55},
			lenSecondArray: 5,
			expected:       []int{1, 5, 6, 7, 11, 14, 15, 21, 25, 29, 37, 55},
		},
		{
			name:           "merge 2 empty arrays",
			firstArray:     []int{},
			lenFirstArray:  0,
			secondArray:    []int{},
			lenSecondArray: 0,
			expected:       []int{},
		},
	}
	if t != nil {
		for _, testCase := range testTable {

			t.Logf("Calling MergeSort(%v, %v)",
				testCase.firstArray, testCase.secondArray)

			MergeSort(testCase.firstArray, testCase.lenFirstArray, testCase.secondArray, testCase.lenSecondArray)

			result := testCase.firstArray
			t.Logf("result: %v", result)
			assert.Equal(t, testCase.expected, result,
				fmt.Sprintf("MergeSort returned %v, expected %v",
					result, testCase.expected))
		}
	}
	if b != nil {
		for _, testCase := range testTable {
			b.Run(testCase.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					MergeSort(testCase.firstArray, testCase.lenFirstArray, testCase.secondArray, testCase.lenSecondArray)
				}
			})
		}
	}
}

// Тест ф-ции на проверку палиндрома без учета регистра и знаков
func TestIsPalindromeUppercase(t *testing.T) {
	TBIsPalindromeUppercase(t, nil)
}

func BenchmarkIsPalindromeUppercase(b *testing.B) {
	TBIsPalindromeUppercase(nil, b)
}

func TBIsPalindromeUppercase(t *testing.T, b *testing.B) {
	testTable := []struct {
		name     string
		text     string
		expected bool
	}{
		{
			name:     "checking palindrome: 7 letters",
			text:     "Rep a Per",
			expected: true,
		},
		{
			name:     "checking palindrome: 9 letters",
			text:     "i am batman",
			expected: false,
		},
		{
			name:     "checking palindrome: 21 letters and punctuations",
			text:     "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			name:     "checking palindrome: 18 letters",
			text:     "Lorem ipsum dolor sit",
			expected: false,
		},
	}
	if t != nil {
		for _, testCase := range testTable {

			result := IsPalindromeUppercase(testCase.text)

			t.Logf("Calling IsPalindromeUppercase(%s), result: %v", testCase.text, result)
			assert.Equal(t, testCase.expected, result,
				fmt.Sprintf("IsPalindromeUppercase returned %v, expected %v",
					result, testCase.expected))
		}
	}
	if b != nil {
		for _, testCase := range testTable {
			b.Run(testCase.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					IsPalindromeUppercase(testCase.text)
				}
			})
		}
	}
}

// Тест ф-ции на проверку палиндрома с возможностью удаления одного символа
func TestIsPalindromeWithoutOneLetter(t *testing.T) {
	TBIsPalindromeWithoutOneLetter(t, nil)
}

func BenchmarkIsPalindromeWithoutOneLetter(b *testing.B) {
	TBIsPalindromeWithoutOneLetter(nil, b)
}

func TBIsPalindromeWithoutOneLetter(t *testing.T, b *testing.B) {
	testTable := []struct {
		name     string
		text     string
		expected bool
	}{
		{
			name:     "checking palindrome without 1 letter: 8 letters",
			text:     "abcdcbad",
			expected: true,
		},
		{
			name:     "checking palindrome without 1 letter: 9 letters",
			text:     "aebcdcbad",
			expected: false,
		},
		{
			name:     "checking palindrome without 1 letter: 17 letters",
			text:     "abcdrzhddhzrdecba",
			expected: true,
		},
		{
			name:     "checking palindrome without 1 letter: 18 letters",
			text:     "abcdrzthddhzrdecba",
			expected: false,
		},
	}
	if t != nil {
		for _, testCase := range testTable {

			result := IsPalindromeWithoutOneLetter(testCase.text)

			t.Logf("Calling IsPalindromeWithoutOneLetter(%s), result: %v", testCase.text, result)
			assert.Equal(t, testCase.expected, result,
				fmt.Sprintf("IsPalindromeWithoutOneLetter returned %v, expected %v",
					result, testCase.expected))
		}
	}
	if b != nil {
		for _, testCase := range testTable {
			b.Run(testCase.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					IsPalindromeWithoutOneLetter(testCase.text)
				}
			})
		}
	}
}

// Тест ф-ции поиска суммы двух чисел из одного среза
func TestTwoSum(t *testing.T) {
	TBTwoSum(t, nil)
}

func BenchmarkTwoSum(b *testing.B) {
	TBTwoSum(nil, b)
}

func TBTwoSum(t *testing.T, b *testing.B) {
	testTable := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "find sum of 2 numbers on array with 4 elements",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "find sum of 2 numbers on array with 7 elements",
			nums:     []int{1, 2, 3, 5, 7, 11, 15},
			target:   13,
			expected: []int{1, 5},
		},
		{
			name:     "find sum of 2 numbers on array with 10 elements",
			nums:     []int{1, 2, 3, 5, 7, 11, 15, 22, 25, 70},
			target:   71,
			expected: []int{0, 9},
		},
	}
	if t != nil {
		for _, testCase := range testTable {

			result := TwoSum(testCase.nums, testCase.target)

			t.Logf("Calling TwoSum(%v), result: %v", testCase.nums, result)
			assert.Equal(t, testCase.expected, result,
				fmt.Sprintf("TwoSum returned %v, expected %v",
					result, testCase.expected))
		}
	}
	if b != nil {
		for _, testCase := range testTable {
			b.Run(testCase.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					TwoSum(testCase.nums, testCase.target)
				}
			})
		}
	}
}

// Тест ф-ции проверки уникальности повторения чисел
func TestUniqueOccurrences(t *testing.T) {
	TBUniqueOccurrences(t, nil)
}

func BenchmarkUniqueOccurrences(b *testing.B) {
	TBUniqueOccurrences(nil, b)
}

func TBUniqueOccurrences(t *testing.T, b *testing.B) {
	testTable := []struct {
		name     string
		arr      []int
		expected bool
	}{
		{
			name:     "checking unique repetition occurrences in array with 6 elements",
			arr:      []int{1, 2, 2, 1, 1, 3},
			expected: true,
		},
		{
			name:     "checking unique repetition occurrences in array with 5 elements",
			arr:      []int{1, 2, 1, 1, 3},
			expected: false,
		},
		{
			name:     "checking unique repetition occurrences in array with 15 elements",
			arr:      []int{1, 2, 2, 1, 1, 3, 6, 4, 6, 4, 6, 4, 4, 6, 4},
			expected: true,
		},
		{
			name:     "checking unique repetition occurrences in array with 14 elements",
			arr:      []int{1, 2, 2, 1, 3, 6, 6, 4, 6, 4, 4, 6, 4},
			expected: false,
		},
	}
	if t != nil {
		for _, testCase := range testTable {

			result := UniqueOccurrences(testCase.arr)

			t.Logf("Calling UniqueOccurrences(%v), result: %v", testCase.arr, result)
			assert.Equal(t, testCase.expected, result,
				fmt.Sprintf("UniqueOccurrences returned %v, expected %v",
					result, testCase.expected))
		}
	}
	if b != nil {
		for _, testCase := range testTable {
			b.Run(testCase.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					UniqueOccurrences(testCase.arr)
				}
			})
		}
	}
}

// Тест ф-ции нахождения главного числа среза
func TestMajorityElement(t *testing.T) {
	TBMajorityElement(t, nil)
}

func BenchmarkMajorityElement(b *testing.B) {
	TBMajorityElement(nil, b)
}

func TBMajorityElement(t *testing.T, b *testing.B) {
	testTable := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "find majority element in array with 3 elements",
			nums:     []int{3, 1, 3},
			expected: 3,
		},
		{
			name:     "find majority element in array with 7 elements",
			nums:     []int{2, 2, 1, 1, 1, 2, 2},
			expected: 2,
		},
		{
			name:     "find majority element in array with 14 elements",
			nums:     []int{5, 5, 1, 6, 1, 5, 5, 2, 5, 5, 5, 5, 2, 8},
			expected: 5,
		},
	}
	if t != nil {
		for _, testCase := range testTable {

			result := MajorityElement(testCase.nums)

			t.Logf("Calling MajorityElement(%v), result: %v", testCase.nums, result)
			assert.Equal(t, testCase.expected, result,
				fmt.Sprintf("MajorityElement returned %v, expected %v",
					result, testCase.expected))
		}
	}
	if b != nil {
		for _, testCase := range testTable {
			b.Run(testCase.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					MajorityElement(testCase.nums)
				}
			})
		}
	}
}

// Тест ф-ции проверки правильности закрытия скобок
func TestParenthesesIsBalanced(t *testing.T) {
	TBParenthesesIsBalanced(t, nil)
}

func BenchmarkParenthesesIsBalanced(b *testing.B) {
	TBParenthesesIsBalanced(nil, b)
}

func TBParenthesesIsBalanced(t *testing.T, b *testing.B) {
	testTable := []struct {
		name     string
		text     string
		expected bool
	}{
		{
			name:     "checking if parentheses are closed correctly. Positive result",
			text:     "({[]})",
			expected: true,
		},
		{
			name:     "checking if parentheses are closed correctly. Negative result",
			text:     "([)]",
			expected: false,
		},
		{
			name:     "checking if parentheses are closed correctly. Empty string",
			text:     "",
			expected: true,
		},
		{
			name:     "checking if parentheses are closed correctly with text inside",
			text:     "({[text]})",
			expected: false,
		},
	}
	if t != nil {
		for _, testCase := range testTable {

			result := IsBalanced(testCase.text)

			t.Logf("Calling IsBalanced(%s), result: %v", testCase.text, result)
			assert.Equal(t, testCase.expected, result,
				fmt.Sprintf("IsBalanced returned %v, expected %v",
					result, testCase.expected))
		}
	}
	if b != nil {
		for _, testCase := range testTable {
			b.Run(testCase.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					IsBalanced(testCase.text)
				}
			})
		}
	}
}

// Тест алгоритма поиска подстроки в строке
func TestNaiveSearch(t *testing.T) {
	TBNaiveSearch(t, nil)
}

func BenchmarkNaiveSearch(b *testing.B) {
	TBNaiveSearch(nil, b)
}

func TBNaiveSearch(t *testing.T, b *testing.B) {
	testTable := []struct {
		name     string
		text     string
		target   string
		expected int
	}{
		{
			name:     "searching for a substring in a string with 1 word",
			text:     "абвгдейка",
			target:   "где",
			expected: 3,
		},
		{
			name:     "searching for a substring in a string with 8 words",
			text:     "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
			target:   "sum",
			expected: 8,
		},
		{
			name:     "searching for a empty substring in a string",
			text:     "Набор любых слов",
			target:   "",
			expected: 0,
		},
		{
			name:     "searching for a substring in a empty string",
			text:     "",
			target:   "target",
			expected: -1,
		},
	}
	if t != nil {
		for _, testCase := range testTable {

			result := NaiveSearch(testCase.text, testCase.target)

			t.Logf("Calling NaiveSearch(%s), result: %v", testCase.text, result)
			assert.Equal(t, testCase.expected, result,
				fmt.Sprintf("NaiveSearch returned %v, expected %v",
					result, testCase.expected))
		}
	}
	if b != nil {
		for _, testCase := range testTable {
			b.Run(testCase.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					NaiveSearch(testCase.text, testCase.target)
				}
			})
		}
	}
}
