package retry

import (
	"math/rand"
	"time"
)

type Err error

type Stop struct {
	Err
}

func Retry(attempts int, sleep time.Duration, f func() error) error {
	if err := f(); err != nil {
		if s, ok := err.(Stop); ok {
			return s.Err
		}

		if attempts > 0 {
			jitter := time.Duration(rand.Int63n(int64(sleep)))
			sleep = sleep + jitter/2

			time.Sleep(sleep)
			attempts--
			return Retry(attempts, 2*sleep, f)
		}
		return err
	}
	return nil
}
