package nux

import (
	"log"
	"testing"
	"time"
)

func TestNtpTime(t *testing.T) {
	orgTime := time.Now()
	log.Println("Begin")
	serverReciveTime, serverTransmitTime, err := NtpTwoTime("ntp1.aliyun.com", 20)
	if err != nil {
		log.Println(err)
		return
	}
	dstTime := time.Now()

	// https://en.wikipedia.org/wiki/Network_Time_Protocol
	duration := ((serverReciveTime.UnixNano() - orgTime.UnixNano()) + (serverTransmitTime.UnixNano() - dstTime.UnixNano())) / 2

	delta := duration / 1e6 // convert to ms
	log.Println(delta)
}
