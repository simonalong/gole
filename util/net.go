package util

import (
	"errors"
	"net"
	"strings"
)

func GetIntranetIp() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, address := range addrs {
		// 检查地址是否为IPv4类型
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", errors.New("没有找到")
}

func GetPublicIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		panic(err)
		return "", err
	}
	addr := conn.LocalAddr().(*net.UDPAddr)
	return strings.Split(addr.String(), ":")[0], nil
}

func GetMacAddress() ([]string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	var macAddresses []string
	for _, iface := range interfaces {
		mac := iface.HardwareAddr.String()
		if mac != "" {
			macAddresses = append(macAddresses, mac)
		}
	}
	return macAddresses, nil
}
