package net

import (
	"net"
	"time"
)

// IpPortAvailable 判断网络是否可达
// ipAndPort：格式：{ip}:{port}
func IpPortAvailable(ipAndPort string) bool {
	conn, err := net.DialTimeout("tcp", ipAndPort, 2*time.Second)
	if err != nil {
		return false
	} else {
		if conn != nil {
			err := conn.Close()
			if err != nil {
				return false
			}
			return true
		} else {
			return false
		}
	}
}
