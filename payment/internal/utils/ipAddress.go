package utils

import (
	"fmt"
	"github.com/davitdarsalia/payment/internal/types"
	"net"
	"os"
)

func IpAddress() (result types.IpV16) {
	host, _ := os.Hostname()
	address, _ := net.LookupIP(host)

	for _, a := range address {
		if ipv16 := a.To16(); ipv16 != nil {
			result = fmt.Sprintf("%s ", ipv16)
		}
	}

	return result
}
