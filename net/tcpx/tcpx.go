package tcpx

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	Hosts                []string
	GlobalTimeout        time.Duration
	TcpConnectionTimeout time.Duration
	WaitBefore           time.Duration
	WaitAfter            time.Duration
	WaitSleepInterval    time.Duration
}

func WaitHosts() {
	config := ConfigFromEnv()
	Wait(config, func() { os.Exit(1) })
}

func ConfigFromEnv() *Config {
	hosts := os.Getenv("WAIT_HOSTS")
	hosts = strings.ReplaceAll(hosts, ",", " ")
	return &Config{
		Hosts:                strings.Split(hosts, ","),
		GlobalTimeout:        envTimeParse("WAIT_TIMEOUT", time.Second*30),
		TcpConnectionTimeout: envTimeParse("WAIT_HOST_CONNECT_TIMEOUT", time.Second*5),
		WaitBefore:           envTimeParse("WAIT_BEFORE", 0),
		WaitAfter:            envTimeParse("WAIT_AFTER", 0),
		WaitSleepInterval:    envTimeParse("WAIT_SLEEP_INTERVAL", time.Second),
	}
}

func IsPortReachableWithTimeout(host string, tpcTimeout time.Duration) bool {
	conn, err := net.DialTimeout("tcp", host, tpcTimeout)
	if err != nil {
		return false
	}
	_ = conn.Close()
	return true
}

func Wait(config *Config, onTimeout func()) {
	if config.WaitBefore > 0 {
		time.Sleep(config.WaitBefore)
	}
	afterTime := time.After(config.GlobalTimeout)

loopLabel:
	for _, host := range config.Hosts {
		for {
			fmt.Printf("[%s] Waiting for host: %s\n", time.Now(), host)
			reach := IsPortReachableWithTimeout(host, config.TcpConnectionTimeout)
			if reach {
				fmt.Printf("[%s] Host is ready: %s\n", time.Now(), host)
				break
			}
			select {
			case <-afterTime:
				fmt.Printf("[%s] Wait for the host to time out: %s\n", time.Now(), host)
				onTimeout()
				break loopLabel
			default:
				time.Sleep(config.WaitSleepInterval)
			}
		}
	}

	if config.WaitAfter > 0 {
		time.Sleep(config.WaitAfter)
	}
}

func envTimeParse(key string, defaultValue time.Duration) time.Duration {
	timeoutEnv := os.Getenv(key)
	if timeoutEnv != "" {
		i, err := strconv.ParseInt(timeoutEnv, 10, 64)
		if err == nil {
			defaultValue = time.Second * time.Duration(i)
		}
	}
	return defaultValue
}
