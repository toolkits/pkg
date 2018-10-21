package str

import (
	"hash/crc32"
	"math/rand"
	"os"
	"time"
)

func init() {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	rand.Seed(time.Now().UnixNano() + int64(os.Getpid()+os.Getppid()) + int64(crc32.ChecksumIEEE([]byte(name))))
}
