package utils

import (
	"fmt"
	"net"
	"os"
)

func IpAddress() (result string) {
	host, _ := os.Hostname()
	address, _ := net.LookupIP(host)

	for _, a := range address {
		if ipv16 := a.To16(); ipv16 != nil {
			result = fmt.Sprintf("%s ", ipv16)
		}
	}

	return result
}
