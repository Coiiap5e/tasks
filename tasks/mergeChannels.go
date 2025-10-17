package tasks

import (
	"context"
	"sync"
)

type Telemetry struct {
	SensorData string
}

func MergeTelemetry(ctx context.Context, sensors ...<-chan Telemetry) <-chan Telemetry {
	var wg sync.WaitGroup
	mergedChannels := make(chan Telemetry)
	wg.Add(len(sensors))

	for _, sensor := range sensors {
		go func(dataChan <-chan Telemetry) {
			defer wg.Done()

			for {
				select {
				case data, ok := <-dataChan:
					if !ok {
						return
					}
					select {
					case mergedChannels <- data:
					case <-ctx.Done():
						return
					}
				case <-ctx.Done():
					return
				}
			}
		}(sensor)
	}

	go func() {
		wg.Wait()
		close(mergedChannels)
	}()

	return mergedChannels
}
