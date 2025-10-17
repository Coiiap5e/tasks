package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Тест поиска числа в отсортированном срезе

func TestBinarySearch(t *testing.T) {
	// Given
	testTable := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "6 elements in array - found",
			nums:     []int{-1, 0, 3, 5, 9, 12},
			target:   5,
			expected: 3,
		},
		{
			name:     "6 elements in array - not found",
			nums:     []int{-1, 0, 3, 5, 9, 12},
			target:   11,
			expected: -1,
		},
		{
			name:     "11 elements in array - found",
			nums:     []int{-1, 0, 3, 5, 9, 12, 15, 17, 32, 45, 51},
			target:   32,
			expected: 8,
		},
		{
			name:     "11 elements in array - found",
			nums:     []int{-1, 0, 3, 5, 9, 12, 15, 17, 32, 45, 51},
			target:   44,
			expected: -1,
		},
	}

	for _, testCase := range testTable {
		//When
		result := BinarySearch(testCase.nums, testCase.target)

		//Then
		t.Logf("calling BinarySearch(%v), result: %d", testCase.nums, result)
		assert.Equal(t, testCase.expected, result,
			fmt.Sprintf("BinarySearch returned %d, expected %d",
				result, testCase.expected))
	}

}

// Тест поиска самой первой "плохой" версии (плохая версия идет с 6й включительно)
func TestFirstBadVersion(t *testing.T) {
	// Given
	testTable := []struct {
		name     string
		versions []int
		expected int
	}{
		{
			name:     "6 elements in array",
			versions: []int{0, 3, 5, 9, 12, 16},
			expected: 3,
		},
		{
			name:     "10 elements in array",
			versions: []int{0, 3, 4, 5, 12, 16, 17, 32, 45, 51},
			expected: 4,
		},
		{
			name:     "10 elements in array",
			versions: []int{0, 1, 2, 3, 4, 5, 6, 9, 12, 16},
			expected: 6,
		},
	}

	for _, testCase := range testTable {
		//When
		result := FirstBadVersion(testCase.versions)

		//Then
		t.Logf("calling FirstBadVersion(%v), result: %d", testCase.versions, result)
		assert.Equal(t, testCase.expected, result,
			fmt.Sprintf("FirstBadVersion returned %d, expected %d",
				result, testCase.expected))
	}

}

// Тест нахождения квадратного корня от числа с округлением в меньшую сторону
func TestRoundedSqrt(t *testing.T) {
	// Given
	testTable := []struct {
		name     string
		number   int
		expected int
	}{
		{
			name:     "Without rounding. Number: 9",
			number:   9,
			expected: 3,
		},
		{
			name:     "With rounding. Number: 8",
			number:   8,
			expected: 2,
		},
		{
			name:     "With rounding. Number: 31",
			number:   31,
			expected: 5,
		},
	}

	for _, testCase := range testTable {
		//When
		result := RoundedSqrt(testCase.number)

		//Then
		t.Logf("calling RoundedSqrt(%d), result: %d", testCase.number, result)
		assert.Equal(t, testCase.expected, result,
			fmt.Sprintf("RoundedSqrt returned %d, expected %d",
				result, testCase.expected))
	}

}
