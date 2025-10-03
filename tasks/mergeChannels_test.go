package tasks

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeTelemetry(t *testing.T) {
	TBMergeTelemetry(t, nil)
}

func BenchmarkMergeTelemetry(b *testing.B) {
	TBMergeTelemetry(nil, b)
}

func TBMergeTelemetry(t *testing.T, b *testing.B) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	testTable := []struct {
		name      string
		dataArray [][]string
		expected  bool
	}{
		{
			name: "three channels with 3 data each",
			dataArray: [][]string{
				{"data1", "data2", "data3"},
				{"data4", "data5", "data6"},
				{"data7", "data8", "data9"},
			},
			expected: true,
		},
		{
			name: "five channels with 6 data each",
			dataArray: [][]string{
				{"data1", "data2", "data3", "data4", "data5", "data6"},
				{"data7", "data8", "data9", "data10", "data11", "data12"},
				{"data13", "data14", "data15", "data16", "data17", "data18"},
				{"data19", "data20", "data21", "data22", "data23", "data24"},
				{"data25", "data26", "data27", "data28", "data29", "data29"},
			},
			expected: true,
		},
		{
			name: "two channels with 6 data each",
			dataArray: [][]string{
				{"data1", "data2", "data3", "data4", "data5", "data6"},
				{"data7", "data8", "data9", "data10", "data11", "data12"},
			},
			expected: true,
		},
	}

	if t != nil {
		for _, testCase := range testTable {
			channelsArray := make([]chan Telemetry, len(testCase.dataArray))

			for i := 0; i < len(testCase.dataArray); i++ {
				channelsArray[i] = make(chan Telemetry, len(testCase.dataArray[i]))
			}

			var wg sync.WaitGroup
			wg.Add(len(channelsArray))

			for i, channel := range channelsArray {
				go func(index int, channel chan Telemetry) {
					defer wg.Done()
					defer close(channel)

					for _, value := range testCase.dataArray[index] {
						channel <- Telemetry{SensorData: value}
					}
				}(i, channel)
			}

			onlyReadChannels := make([]<-chan Telemetry, len(testCase.dataArray))

			for i, channel := range channelsArray {
				onlyReadChannels[i] = channel
			}

			lenSlice := 0
			for _, slice := range testCase.dataArray {
				lenSlice += len(slice)
			}
			dataResult := make(map[Telemetry]int, lenSlice)
			for _, data := range testCase.dataArray {
				for _, value := range data {
					key := Telemetry{SensorData: value}
					dataResult[key]++
				}
			}
			
			// вызов ключевой ф-ции
			mergedChan := MergeTelemetry(ctx, onlyReadChannels...)

			for data := range mergedChan {
				dataResult[data]--
			}
			result := true
			for _, data := range dataResult {
				if data != 0 {
					result = false
				}
			}

			t.Logf("Merge %v channels with func MergeTelemetry(), result: %v",
				len(testCase.dataArray), result)
			assert.Equal(t, testCase.expected, result,
				fmt.Sprintf("MergeTelemetry() returned %v, expected %v",
					result, testCase.expected))

			wg.Wait()

			t.Logf("Checking if %v channels are closed", len(testCase.dataArray))
			for i := 0; i < len(testCase.dataArray); i++ {
				select {
				case _, ok := <-channelsArray[i]:
					if ok {
						t.Errorf("Channel %v still open", i)
					} else {
						t.Logf("Channel %v are closed", i)
					}
				default:
					t.Errorf("Channel %v still open and empty", i)
				}
			}
			select {
			case _, ok := <-mergedChan:
				if ok {
					t.Error("Merged channel still open")
				} else {
					t.Log("Merged channel are closed")
				}
			default:
				t.Error("Merged channel still open and empty")
			}

		}
	}

	if b != nil {
		for _, testCase := range testTable {
			channelsArray := make([]chan Telemetry, len(testCase.dataArray))

			for i := 0; i < len(testCase.dataArray); i++ {
				channelsArray[i] = make(chan Telemetry, len(testCase.dataArray[i]))
			}

			for i, channel := range channelsArray {
				go func(index int, channel chan Telemetry) {
					defer close(channel)

					for _, value := range testCase.dataArray[index] {
						channel <- Telemetry{SensorData: value}
					}
				}(i, channel)
			}

			onlyReadChannels := make([]<-chan Telemetry, len(testCase.dataArray))

			for i, channel := range channelsArray {
				onlyReadChannels[i] = channel
			}

			b.Run(testCase.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					MergeTelemetry(ctx, onlyReadChannels...)
				}
			})
		}
	}
}
