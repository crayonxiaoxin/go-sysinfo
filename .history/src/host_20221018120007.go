package src

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

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
		{"Uptime", strconv.FormatUint(info.Uptime, 10)},
		{"BootTime", strconv.FormatUint(info.BootTime, 10)},
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

	fmt.Printf("info.BootTime: %v\n", info.BootTime)
	fmt.Printf("info.BootTime: %v\n", time.Unix(int64(info.BootTime), 0))
	fmt.Printf("info.BootTime: %v\n", time.Unix(int64(info.BootTime), 0).Format("2006-01-02 15:04:05"))

	table.Render()
}
