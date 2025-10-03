package tasks

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Тест счетчика, реализуемый через атомики и который работает с множества горутин
func TestRunAtomicCounter(t *testing.T) {
	TBRunAtomicCounter(t, nil)
}

func BenchmarkRunAtomicCounter(b *testing.B) {
	TBRunAtomicCounter(nil, b)
}

func TBRunAtomicCounter(t *testing.T, b *testing.B) {
	testTable := []struct {
		name           string
		incCount       int
		goroutineCount int
		expected       int64
	}{
		{
			name:           "one goroutine in which count 1 times",
			incCount:       1,
			goroutineCount: 1,
			expected:       1,
		},
		{
			name:           "loop 1000 goroutine in which loop a counter 10000 times",
			incCount:       10000,
			goroutineCount: 1000,
			expected:       10000000,
		},
		{
			name:           "loop 5000 goroutine in which loop a counter 1000 times",
			incCount:       1000,
			goroutineCount: 5000,
			expected:       5000000,
		},
		{
			name:           "loop 100 goroutine in which loop a counter 400000 times",
			incCount:       400000,
			goroutineCount: 100,
			expected:       40000000,
		},
	}

	if t != nil {
		for _, testCase := range testTable {

			result := RunAtomicCounter(testCase.incCount, testCase.goroutineCount)

			t.Logf("Calling RunAtomicCounter(%v, %v), result: %v",
				testCase.incCount, testCase.goroutineCount, result)
			assert.Equal(t, testCase.expected, result,
				fmt.Sprintf("RunAtomicCounter returned %v, expected %v",
					result, testCase.expected))
		}
	}

	if b != nil {
		for _, testCase := range testTable {
			b.Run(testCase.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					RunAtomicCounter(testCase.incCount, testCase.goroutineCount)
				}
			})
		}
	}
}

// Тест счетчика, реализуемый через мьютексы и который работает с множества горутин
func TestRunMutexCounter(t *testing.T) {
	TBRunMutexCounter(t, nil)
}

func BenchmarkRunMutexCounter(b *testing.B) {
	TBRunMutexCounter(nil, b)
}

func TBRunMutexCounter(t *testing.T, b *testing.B) {
	testTable := []struct {
		name           string
		incCount       int
		goroutineCount int
		expected       int
	}{
		{
			name:           "one goroutine in which count 1 times",
			incCount:       1,
			goroutineCount: 1,
			expected:       1,
		},
		{
			name:           "loop 1000 goroutine in which loop a counter 10000 times",
			incCount:       10000,
			goroutineCount: 1000,
			expected:       10000000,
		},
		{
			name:           "loop 5000 goroutine in which loop a counter 1000 times",
			incCount:       1000,
			goroutineCount: 5000,
			expected:       5000000,
		},
		{
			name:           "loop 100 goroutine in which loop a counter 400000 times",
			incCount:       400000,
			goroutineCount: 100,
			expected:       40000000,
		},
	}

	if t != nil {
		for _, testCase := range testTable {

			result := RunMutexCounter(testCase.incCount, testCase.goroutineCount)

			t.Logf("Calling RunMutexCounter(%v, %v), result: %v",
				testCase.incCount, testCase.goroutineCount, result)
			assert.Equal(t, testCase.expected, result,
				fmt.Sprintf("RunMutexCounter returned %v, expected %v",
					result, testCase.expected))
		}
	}

	if b != nil {
		for _, testCase := range testTable {
			b.Run(testCase.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					RunMutexCounter(testCase.incCount, testCase.goroutineCount)
				}
			})
		}
	}
}
