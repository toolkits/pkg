package watcher

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Watcher struct {
	Pid int
	Exe string
	Dur time.Duration
	Dev bool
}

func NewWatcher(exe string, dur time.Duration) *Watcher {
	return &Watcher{
		Exe: exe,
		Dur: dur,
	}
}

func (w *Watcher) Start(f func()) {
	// clean deleted process
	fs, err := ioutil.ReadDir("/proc")
	if err != nil {
		if w.Dev {
			log.Println("ERR: cannot read /proc:", err)
		}
		return
	}

	sz := len(fs)
	for i := 0; i < sz; i++ {
		if !fs[i].IsDir() {
			continue
		}

		name := fs[i].Name()
		pid, err := strconv.Atoi(name)
		if err != nil {
			continue
		}

		exe := fmt.Sprintf("/proc/%d/exe", pid)
		if !IsExist(exe) {
			continue
		}

		target, err := os.Readlink(exe)
		if err == nil && strings.Contains(target, w.Exe) && strings.Contains(target, "deleted") {
			proc, err := os.FindProcess(pid)
			if err != nil {
				if w.Dev {
					log.Printf("ERR: cannot find process[pid:%d]: %v", pid, err)
				}
				continue
			}

			err = proc.Kill()
			if err != nil && w.Dev {
				log.Printf("ERR: cannot kill process[pid:%d]: %v", pid, err)
			}
		}
	}

	for {
		time.Sleep(w.Dur)
		w.check(f)
	}
}

func (w *Watcher) check(f func()) {
	defer func() {
		if err := recover(); err != nil {
			if w.Dev {
				log.Println("PANIC:", err)
			}
			return
		}
	}()

	if w.Pid > 0 {
		target, err := os.Readlink(fmt.Sprintf("/proc/%d/exe", w.Pid))
		if err == nil && target == w.Exe {
			return
		}
	}

	fs, err := ioutil.ReadDir("/proc")
	if err != nil {
		if w.Dev {
			log.Println("ERR: cannot read /proc:", err)
		}
		return
	}

	sz := len(fs)
	for i := 0; i < sz; i++ {
		if !fs[i].IsDir() {
			continue
		}

		name := fs[i].Name()
		pid, err := strconv.Atoi(name)
		if err != nil {
			continue
		}

		exe := fmt.Sprintf("/proc/%d/exe", pid)
		if !IsExist(exe) {
			continue
		}

		target, err := os.Readlink(exe)
		if err == nil && target == w.Exe {
			w.Pid = pid
			return
		}
	}

	// w.Exe not found
	f()
}

func IsExist(fp string) bool {
	_, err := os.Stat(fp)
	return err == nil || os.IsExist(err)
}
