package aggregate

import (
	"context"
	"log"
)

// Implements ResultListener interface
type LoggingResultListener struct {
	resultListener
	source <-chan *Sample
}

func NewLoggingResultListener() ResultListener {
	ch := make(chan *Sample, 32)
	return &LoggingResultListener{
		source: ch,
		resultListener: resultListener{
			sink: ch,
		},
	}
}

func (rl *LoggingResultListener) handle(s *Sample) {
	log.Println(s)
	ReleaseSample(s)
}

func (rl *LoggingResultListener) Start(ctx context.Context) error {
loop:
	for {
		select {
		case r := <-rl.source:
			rl.handle(r)
		case <-ctx.Done():
			// Context is done, but we should read all data from source
			for {
				select {
				case r := <-rl.source:
					rl.handle(r)
				default:
					break loop
				}
			}
		}
	}
	return nil
}