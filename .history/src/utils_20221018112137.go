package src

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

const TOKEN_IPINFO = "5a1445e6b70829"

func SizeUnit(size uint64) string {
	bytes := float64(size)
	kb := float64(1024)
	mb := 1024 * kb
	gb := 1024 * mb
	if bytes >= gb {
		return strconv.FormatFloat(bytes/gb, 'f', 2, 64) + " GB"
	} else if bytes >= mb {
		return strconv.FormatFloat(bytes/mb, 'f', 2, 64) + " MB"
	} else if bytes >= kb {
		return strconv.FormatFloat(bytes/kb, 'f', 2, 64) + " KB"
	} else if bytes > 0 {
		return strconv.FormatFloat(bytes, 'f', 2, 64) + " Bytes"
	} else {
		return "0"
	}
}

func LocalIP() string {
	con, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return ""
	}
	defer con.Close()
	ip := con.LocalAddr().String()
	ip = strings.Split(ip, ":")[0]
	return ip
}