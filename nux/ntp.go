// Copyright 2015 Brett Vickers.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package ntp provides a simple mechanism for querying the current time
// from a remote NTP server.  This package only supports NTP client mode
// behavior and version 4 of the NTP protocol.  See RFC 5905.
// Approach inspired by go-nuts post by Michael Hofmann:
// https://groups.google.com/forum/?fromgroups#!topic/golang-nuts/FlcdMU5fkLQ
package nux

import (
	"encoding/binary"
	"net"
	"time"
)

type mode byte

const (
	reserved mode = 0 + iota
	symmetricActive
	symmetricPassive
	client
	server
	broadcast
	controlMessage
	reservedPrivate
)

type ntpTime struct {
	Seconds  uint32
	Fraction uint32
}

func (t ntpTime) UTC() time.Time {
	nsec := uint64(t.Seconds)*1e9 + (uint64(t.Fraction) * 1e9 >> 32)
	return time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC).Add(time.Duration(nsec))
}

type msg struct {
	LiVnMode       byte // Leap Indicator (2) + Version (3) + Mode (3)
	Stratum        byte
	Poll           byte
	Precision      byte
	RootDelay      uint32
	RootDispersion uint32
	ReferenceId    uint32
	ReferenceTime  ntpTime
	OriginTime     ntpTime
	ReceiveTime    ntpTime
	TransmitTime   ntpTime
}

// SetVersion sets the NTP protocol version on the message.
func (m *msg) SetVersion(v byte) {
	m.LiVnMode = (m.LiVnMode & 0xc7) | v<<3
}

// SetMode sets the NTP protocol mode on the message.
func (m *msg) SetMode(md mode) {
	m.LiVnMode = (m.LiVnMode & 0xf8) | byte(md)
}

// Time returns the "receive time" from the remote NTP server
// specifed as host.  NTP client mode is used.
func getTwoTime(host string, version byte, timeout int64) (time.Time, time.Time, error) {
	if version < 2 || version > 4 {
		panic("ntp: invalid version number")
	}

	raddr, err := net.ResolveUDPAddr("udp", host+":123")
	if err != nil {
		return time.Now(), time.Now(), err
	}

	con, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		return time.Now(), time.Now(), err
	}
	defer con.Close()
	con.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))

	m := new(msg)
	m.SetMode(client)
	m.SetVersion(version)

	err = binary.Write(con, binary.BigEndian, m)
	if err != nil {
		return time.Now(), time.Now(), err
	}

	err = binary.Read(con, binary.BigEndian, m)
	if err != nil {
		return time.Now(), time.Now(), err
	}

	t := m.ReceiveTime.UTC().Local()
	transmitTime := m.TransmitTime.UTC().Local()
	return t, transmitTime, nil
}

func getTime(host string, version byte, timeout int64) (time.Time, error) {
	if version < 2 || version > 4 {
		panic("ntp: invalid version number")
	}

	raddr, err := net.ResolveUDPAddr("udp", host+":123")
	if err != nil {
		return time.Now(), err
	}

	con, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		return time.Now(), err
	}
	defer con.Close()
	con.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
	m := new(msg)
	m.SetMode(client)
	m.SetVersion(version)

	err = binary.Write(con, binary.BigEndian, m)
	if err != nil {
		return time.Now(), err
	}

	err = binary.Read(con, binary.BigEndian, m)
	if err != nil {
		return time.Now(), err
	}

	t := m.ReceiveTime.UTC().Local()
	return t, nil
}

func getTimeout(timeout []int64) int64 {
	// If the timeout parameter is not provided or the timeout parameter is greater than 0
	if len(timeout) != 1 {
		return 5
	} else {
		return timeout[0]
	}
}

// TimeV returns the "receive time" from the remote NTP server
// specifed as host.  Use the NTP client mode with the requested
// version number (2, 3, or 4).
func NtpTimeV(host string, version byte, timeout ...int64) (time.Time, error) {
	return getTime(host, version, getTimeout(timeout))
}

// Time returns the "receive time" from the remote NTP server
// specifed as host.  NTP client mode version 4 is used.
func NtpTime(host string, timeout ...int64) (time.Time, error) {
	return getTime(host, 4, getTimeout(timeout))
}

// Time returns the "receive time" from the remote NTP server
// specifed as host.  NTP client mode version 4 is used.
func NtpTwoTime(host string, timeout ...int64) (time.Time, time.Time, error) {
	return getTwoTime(host, 4, getTimeout(timeout))
}
