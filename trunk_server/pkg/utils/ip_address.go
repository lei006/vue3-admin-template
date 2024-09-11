package utils

import (
	"fmt"
	"net"
)

func GetAllLocalIPs() ([]string, error) {
	var ips []string

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for _, addr := range addrs {
		// 检查是否为 IP 地址
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String()) // IPv4 地址
			} else {
				ips = append(ips, ipNet.IP.String()) // IPv6 地址
			}
		}
	}

	if len(ips) == 0 {
		return nil, fmt.Errorf("未找到有效的本地 IP 地址")
	}

	return ips, nil
}
