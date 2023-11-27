package main

import (
	"fmt"
	"net"
	"time"
)

func Check(domain string, port string) string {
	addr := domain + ":" + port
	timeout := time.Duration(5 + time.Second)

	conn, err := net.DialTimeout("tcp", addr, timeout)

	var status string
	if err != nil {
		status = fmt.Sprintf("[DOWN] %v is unreachable, \n Error: %v", domain, err)
	} else {
		status = fmt.Sprintf("[UP] %v is reachable, \n From: %v\n To: %v", domain, conn.LocalAddr(), conn.RemoteAddr())
	}

	return status
}
