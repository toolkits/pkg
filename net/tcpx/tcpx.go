package tcpx

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func WaitHosts() {
	hosts := os.Getenv("WAIT_HOSTS")
	if len(hosts) == 0 {
		return
	}

	hosts = strings.ReplaceAll(hosts, ",", " ")
	array := strings.Fields(hosts)

	for _, host := range array {
		waitHost(host)
	}
}

func waitHost(host string) {
	for {
		fmt.Printf("[%s] Waiting for host: %s\n", time.Now(), host)

		conn, err := net.DialTimeout("tcp", host, time.Second)
		if err == nil {
			conn.Close()
			fmt.Printf("[%s] Host is ready: %s\n", time.Now(), host)
			return
		}

		time.Sleep(time.Second)
	}
}
