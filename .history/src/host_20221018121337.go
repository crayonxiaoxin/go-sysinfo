package src

import (
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/fatih/color"
	"github.com/ipinfo/go/v2/ipinfo"
	"github.com/shirou/gopsutil/v3/host"
)

func HostInfo() {
	info, err := host.Info()
	if err != nil {
		log.Fatal(err)
		return
	}

	c := color.New(color.FgHiWhite)
	c = c.Add(color.BgBlue)
	c = c.Add(color.Bold)
	c.Print("           Host           ")
	fmt.Println("")

	table := KeyValueTable()

	table.AppendBulk([][]string{
		{"Hostname", info.Hostname},
		// {"Uptime", strconv.FormatUint(info.Uptime, 10)},
		{"Uptime", SecFmt(int64(info.Uptime))},
		{"BootTime", SecFmt(int64(info.BootTime))},
		{"Procs", strconv.FormatUint(info.Procs, 10)},
		{"OS", info.OS},
		{"Platform", info.Platform},
		{"PlatformFamily", info.PlatformFamily},
		{"PlatformVersion", info.PlatformVersion},
		{"KernelArch", info.KernelArch},
		{"KernelVersion", info.KernelVersion},
		{"VirtualizationSystem", info.VirtualizationSystem},
		{"VirtualizationRole", info.VirtualizationRole},
		{"LocalIP", LocalIP()},
		{"", ""},
	})

	client := ipinfo.NewClient(nil, nil, TOKEN_IPINFO)
	ip, err := client.GetIPAddr()
	if err == nil {
		info, err := client.GetIPInfo(net.ParseIP(ip))
		if err != nil {
			fmt.Printf("err: %v\n", err)
		} else {
			table.AppendBulk([][]string{
				{"IP", info.IP.String()},
				{"City", info.City},
				{"Region", info.Region},
				{"Country", info.Country},
				{"CountryName", info.CountryName},
				{"Location", info.Location},
				{"Org", info.Org},
				{"Timezone", info.Timezone},
			})
		}
	}

	min := uint64(60)
	hour := 60 * min
	day := 24 * hour
	uptime := info.Uptime
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
	} else if uptime >= hour {
		hours := uptime / hour
		a := uptime % hour
		minutes := a / min
		seconds := a % min
		fmt.Printf("hours: %v\n", hours)
		fmt.Printf("minutes: %v\n", minutes)
		fmt.Printf("seconds: %v\n", seconds)
	} else if uptime >= min {
		minutes := uptime / min
		a := uptime % min
		seconds := a % min
		fmt.Printf("minutes: %v\n", minutes)
		fmt.Printf("seconds: %v\n", seconds)
	} else {
		fmt.Printf("seconds: %v\n", uptime)
	}

	table.Render()
}
