package main

import (
	"fmt"
	"net"
	"time"
)

func Check(destination string, port string) string {
	Address := destination + ":" + port
	timeout := time.Duration(5 * time.Second)

	con, err := net.DialTimeout("tcp", Address, timeout)

	var status string

	if err != nil {
		status = fmt.Sprintf("[DOWN] %v is unreachable, \n Error: %v", destination, err)
	} else {
		status = fmt.Sprintf("[UP]  %v is reachable,\n From: %v\n To: %v", destination, con.LocalAddr(), con.RemoteAddr())
	}
	return status

}
