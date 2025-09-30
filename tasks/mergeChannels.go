package tasks

import (
	"context"
	"sync"
)

type Telemetry struct {
	sensorData string
}

func MergeTelemetry(ctx context.Context, sensors ...<-chan Telemetry) <-chan Telemetry {
	var wg sync.WaitGroup
	mergedChannels := make(chan Telemetry)
	wg.Add(len(sensors))
	output := func(dataChan <-chan Telemetry) {
		for data := range dataChan {
			mergedChannels <- data
		}
		wg.Done()
	}

	for _, sensor := range sensors {
		go output(sensor)
	}

	go func() {
		wg.Wait()
		close(mergedChannels)
	}()

	return mergedChannels
}
