package runner

import (
	"hash/crc32"
	"log"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/toolkits/pkg/file"
)

var (
	Hostname  string
	Cwd       string
	CPUNumber int
)

func getCPUNumber() int {
	if value, exist := os.LookupEnv("GOMAXPROCS"); exist {
		if cpunum, err := strconv.Atoi(value); err != nil {
			log.Fatalf("[F] cannot convert env:GOMAXPROCS[%s] %v\n", value, err)
		} else {
			return cpunum
		}
	}

	return runtime.NumCPU()
}

func Init() {
	CPUNumber = getCPUNumber()
	runtime.GOMAXPROCS(CPUNumber)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	var err error
	Hostname, err = os.Hostname()
	if err != nil {
		log.Fatalln("[F] cannot get hostname")
	}

	Cwd = file.SelfDir()

	rand.Seed(time.Now().UnixNano() + int64(os.Getpid()+os.Getppid()) + int64(crc32.ChecksumIEEE([]byte(Hostname))))
}
