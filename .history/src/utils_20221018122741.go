package src

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
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

func KeyValueTable() *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetTablePadding("\t")
	table.SetAutoWrapText(true)
	table.SetHeader([]string{"Key", "Value"})
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.FgRedColor, tablewriter.Bold},
		tablewriter.Colors{tablewriter.FgGreenColor, tablewriter.Bold},
	)
	table.SetColumnAlignment([]int{tablewriter.ALIGN_LEFT, tablewriter.ALIGN_RIGHT})
	table.SetColumnColor(
		tablewriter.Colors{tablewriter.FgRedColor},
		tablewriter.Colors{tablewriter.FgGreenColor},
	)
	return table
}

func Sec2Date(sec int64) string {
	return time.Unix(sec, 0).Format("2006-01-02 15:04:05")
}

func Sec2Diff(sec uint64) string {
	min := uint64(60)
	hour := 60 * min
	day := 24 * hour
	uptime := sec
	if uptime >= day {
		days := uptime / day
		a := uptime % day
		hours := a / hour
		a = a % hour
		minutes := a / min
		seconds := a % min
		fmt.Printf("days: %v\n", days)
		fmt.Printf("hours: %v\n", hours)
		fmt.Printf("minutes: %v\n", minutes)
		fmt.Printf("seconds: %v\n", seconds)
		return fmt.Sprintf("%v days %vh %vm %vs", days, hours, minutes, seconds)
	} else if uptime >= hour {
		hours := uptime / hour
		a := uptime % hour
		minutes := a / min
		seconds := a % min
		fmt.Printf("hours: %v\n", hours)
		fmt.Printf("minutes: %v\n", minutes)
		fmt.Printf("seconds: %v\n", seconds)
		return fmt.Sprintf("%v h%v m%vs", hours, minutes, seconds)
	} else if uptime >= min {
		minutes := uptime / min
		a := uptime % min
		seconds := a % min
		fmt.Printf("minutes: %v\n", minutes)
		fmt.Printf("seconds: %v\n", seconds)
		return fmt.Sprintf("%vm %vs", minutes, seconds)
	} else {
		fmt.Printf("seconds: %v\n", uptime)
		return fmt.Sprintf("%vs", uptime)
	}
}
