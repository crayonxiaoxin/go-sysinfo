package src

import (
	"fmt"
	"log"
	"net"

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

	table := KeyValueTable(nil)

	table.AppendBulk([][]string{
		{"Hostname", info.Hostname},
		{"Uptime", Sec2Diff(info.Uptime)},
		{"BootTime", Sec2Date(int64(info.BootTime))},
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

	table.Render()
}
