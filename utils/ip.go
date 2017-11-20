package utils

import (
	"net"
	"fmt"
	"os"
	"strings"
)

func GetIp() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}

		}
	}
	return ""
}

func GetMac() string {
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, inter := range interfaces {
			if inter.Name == "以太网" {
				return strings.ToLower(fmt.Sprint(inter.HardwareAddr))
			}
		}
	}
	return ""
}
