// +build !plan9

package tools

import (
	"os"
	"os/signal"
	"syscall"
)

func OnInterrupt(fn func()) {
	signalChan := make(chan os.Signal, 1)
	signal.Ignore(syscall.SIGHUP)
	signal.Notify(signalChan,
		os.Interrupt,
		os.Kill,
		syscall.SIGALRM,
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	go func() {
		for _ = range signalChan {
			fn()
			os.Exit(0)
		}
	}()
}
