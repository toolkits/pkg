package retry

import (
	"fmt"
	"testing"
	"time"
)

func TestRetry(t *testing.T) {
	fn := func() error {
		return fmt.Errorf("timeout")
	}

	if err := Retry(1, 2*time.Second, fn); err != nil {
		t.Fatal(err)
	}
}

func TestRetryWithFuncArgs(t *testing.T) {
	// closure fn
	fn := func() error {
		name := "tom"

		// args function
		return func(name string) error {
			t.Logf("name: %s", name)

			return fmt.Errorf("timeout")

			// if returns the retry, your need to returns a Stop instance
			// return Stop{fmt.Errorf("stop")}
		}(name)
	}

	if err := Retry(1, 2*time.Second, fn); err != nil {
		t.Fatal(err)
	}
}
