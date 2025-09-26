package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseStrings(t *testing.T) {
	TBReverseStrings(t, nil)
}

func BenchmarkReverseStrings(b *testing.B) {
	TBReverseStrings(nil, b)
}

func TestMergeSort(t *testing.T) {
	TBMergeSort(t, nil)
}

func BenchmarkMergeSort(b *testing.B) {
	TBMergeSort(nil, b)
}

func TestIsPalindromeUppercase(t *testing.T) {
	TBIsPalindromeUppercase(t, nil)
}

func BenchmarkIsPalindromeUppercase(b *testing.B) {
	TBIsPalindromeUppercase(nil, b)
}

func TestIsPalindromeWithoutOneLetter(t *testing.T) {
	TBIsPalindromeWithoutOneLetter(t, nil)
}

func BenchmarkIsPalindromeWithoutOneLetter(b *testing.B) {
	TBIsPalindromeWithoutOneLetter(nil, b)
}

func TestTwoSum(t *testing.T) {
	TBTwoSum(t, nil)
}

func BenchmarkTwoSum(b *testing.B) {
	TBTwoSum(nil, b)
}

func TestUniqueOccurrences(t *testing.T) {
	TBUniqueOccurrences(t, nil)
}

func BenchmarkUniqueOccurrences(b *testing.B) {
	TBUniqueOccurrences(nil, b)
}

func TestMajorityElement(t *testing.T) {
	TBMajorityElement(t, nil)
}

func BenchmarkMajorityElement(b *testing.B) {
	TBMajorityElement(nil, b)
}

func TestParenthesesIsBalanced(t *testing.T) {
	TBParenthesesIsBalanced(t, nil)
}

func BenchmarkParenthesesIsBalanced(b *testing.B) {
	TBParenthesesIsBalanced(nil, b)
}

func TestNaiveSearch(t *testing.T) {
	TBNaiveSearch(t, nil)
}

func BenchmarkNaiveSearch(b *testing.B) {
	TBNaiveSearch(nil, b)
}

func TBReverseStrings(t *testing.T, b *testing.B) {
	testTable := []struct {
		name     string
		text     string
		expected string
	}{
		{
			name:     "short",
			text:     "hello world",
			expected: "dlrow olleh",
		},
		{
			name:     "empty",
			text:     "",
			expected: "",
		},
		{
			name:     "medium",
			text:     "hello world this is a test",
			expected: "tset a si siht dlrow olleh",
		},
		{
			name:     "long",
			text:     "Lorem ipsum dolor sit amet consectetur adipiscing elit",
			expected: "tile gnicsipida rutetcesnoc tema tis rolod muspi meroL",
		},
	}
	if t != nil {
		for _, testCase := range testTable {
			result := ReverseStrings(testCase.text)
			t.Logf("Calling ReverseStrings(%s), result: %s", testCase.text, result)
			assert.Equal(t, result, testCase.expected,
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
			name:           "len 6",
			firstArray:     []int{1, 2, 3, 0, 0, 0},
			lenFirstArray:  3,
			secondArray:    []int{2, 5, 6},
			lenSecondArray: 3,
			expected:       []int{1, 2, 2, 3, 5, 6},
		},
		{
			name:           "len 12",
			firstArray:     []int{1, 5, 7, 11, 15, 21, 25, 0, 0, 0, 0, 0},
			lenFirstArray:  7,
			secondArray:    []int{6, 14, 29, 37, 55},
			lenSecondArray: 5,
			expected:       []int{1, 5, 6, 7, 11, 14, 15, 21, 25, 29, 37, 55},
		},
		{
			name:           "empty",
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
			assert.Equal(t, result, testCase.expected,
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

func TBIsPalindromeUppercase(t *testing.T, b *testing.B) {
	testTable := []struct {
		name     string
		text     string
		expected bool
	}{
		{
			name:     "short true",
			text:     "Rep a Per",
			expected: true,
		},
		{
			name:     "short false",
			text:     "i am batman",
			expected: false,
		},
		{
			name:     "medium true",
			text:     "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			name:     "medium false",
			text:     "Lorem ipsum dolor sit",
			expected: false,
		},
	}
	if t != nil {
		for _, testCase := range testTable {
			result := IsPalindromeUppercase(testCase.text)
			t.Logf("Calling IsPalindromeUppercase(%s), result: %v", testCase.text, result)
			assert.Equal(t, result, testCase.expected,
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

func TBIsPalindromeWithoutOneLetter(t *testing.T, b *testing.B) {
	testTable := []struct {
		name     string
		text     string
		expected bool
	}{
		{
			name:     "short true",
			text:     "abcdcbad",
			expected: true,
		},
		{
			name:     "short false",
			text:     "aebcdcbad",
			expected: false,
		},
		{
			name:     "medium true",
			text:     "abcdrzhddhzrdecba",
			expected: true,
		},
		{
			name:     "medium false",
			text:     "abcdrzthddhzrdecba",
			expected: false,
		},
	}
	if t != nil {
		for _, testCase := range testTable {
			result := IsPalindromeWithoutOneLetter(testCase.text)
			t.Logf("Calling IsPalindromeWithoutOneLetter(%s), result: %v", testCase.text, result)
			assert.Equal(t, result, testCase.expected,
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

func TBTwoSum(t *testing.T, b *testing.B) {
	testTable := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "len 4",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "len 7",
			nums:     []int{1, 2, 3, 5, 7, 11, 15},
			target:   13,
			expected: []int{1, 5},
		},
		{
			name:     "len 10",
			nums:     []int{1, 2, 3, 5, 7, 11, 15, 22, 25, 70},
			target:   71,
			expected: []int{0, 9},
		},
	}
	if t != nil {
		for _, testCase := range testTable {
			result := TwoSum(testCase.nums, testCase.target)
			t.Logf("Calling TwoSum(%v), result: %v", testCase.nums, result)
			assert.Equal(t, result, testCase.expected,
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

func TBUniqueOccurrences(t *testing.T, b *testing.B) {
	testTable := []struct {
		name     string
		arr      []int
		expected bool
	}{
		{
			name:     "short true",
			arr:      []int{1, 2, 2, 1, 1, 3},
			expected: true,
		},
		{
			name:     "short false",
			arr:      []int{1, 2, 1, 1, 3},
			expected: false,
		},
		{
			name:     "medium true",
			arr:      []int{1, 2, 2, 1, 1, 3, 6, 4, 6, 4, 6, 4, 4, 6, 4},
			expected: true,
		},
		{
			name:     "medium false",
			arr:      []int{1, 2, 2, 1, 1, 3, 6, 6, 4, 6, 4, 4, 6, 4},
			expected: false,
		},
	}
	if t != nil {
		for _, testCase := range testTable {
			result := UniqueOccurrences(testCase.arr)
			t.Logf("Calling UniqueOccurrences(%v), result: %v", testCase.arr, result)
			assert.Equal(t, result, testCase.expected,
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

func TBMajorityElement(t *testing.T, b *testing.B) {
	testTable := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "short",
			nums:     []int{3, 1, 3},
			expected: 3,
		},
		{
			name:     "medium",
			nums:     []int{2, 2, 1, 1, 1, 2, 2},
			expected: 2,
		},
		{
			name:     "large",
			nums:     []int{5, 5, 1, 6, 1, 5, 5, 2, 5, 5, 5, 5, 2, 8},
			expected: 5,
		},
	}
	if t != nil {
		for _, testCase := range testTable {
			result := MajorityElement(testCase.nums)
			t.Logf("Calling MajorityElement(%v), result: %v", testCase.nums, result)
			assert.Equal(t, result, testCase.expected,
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

func TBParenthesesIsBalanced(t *testing.T, b *testing.B) {
	testTable := []struct {
		name     string
		text     string
		expected bool
	}{
		{
			name:     "short true",
			text:     "({[]})",
			expected: true,
		},
		{
			name:     "short false",
			text:     "([)]",
			expected: false,
		},
		{
			name:     "empty",
			text:     "",
			expected: true,
		},
		{
			name:     "with text",
			text:     "({[text]})",
			expected: false,
		},
	}
	if t != nil {
		for _, testCase := range testTable {
			result := IsBalanced(testCase.text)
			t.Logf("Calling IsBalanced(%s), result: %v", testCase.text, result)
			assert.Equal(t, result, testCase.expected,
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

func TBNaiveSearch(t *testing.T, b *testing.B) {
	testTable := []struct {
		name     string
		text     string
		target   string
		expected int
	}{
		{
			name:     "short",
			text:     "абвгдейка",
			target:   "где",
			expected: 3,
		},
		{
			name:     "large",
			text:     "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
			target:   "sum",
			expected: 8,
		},
		{
			name:     "empty target",
			text:     "Набор любых слов",
			target:   "",
			expected: 0,
		},
		{
			name:     "empty text",
			text:     "",
			target:   "target",
			expected: -1,
		},
	}
	if t != nil {
		for _, testCase := range testTable {
			result := NaiveSearch(testCase.text, testCase.target)
			t.Logf("Calling NaiveSearch(%s), result: %v", testCase.text, result)
			assert.Equal(t, result, testCase.expected,
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
