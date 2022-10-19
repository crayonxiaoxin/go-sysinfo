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

	table.Append([]string{"Hostname", info.Hostname})
	table.Append([]string{"Uptime", strconv.FormatUint(info.Uptime, 10)})
	table.Append([]string{"BootTime", strconv.FormatUint(info.BootTime, 10)})
	table.Append([]string{"OS", info.OS})
	table.Append([]string{"Platform", info.Platform})
	table.Append([]string{"PlatformFamily", info.PlatformFamily})
	table.Append([]string{"PlatformVersion", info.PlatformVersion})
	table.Append([]string{"KernelArch", info.KernelArch})
	table.Append([]string{"KernelVersion", info.KernelVersion})
	table.Append([]string{"VirtualizationSystem", info.VirtualizationSystem})
	table.Append([]string{"VirtualizationRole", info.VirtualizationRole})
	table.Append([]string{"LocalIP", LocalIP()})

	client := ipinfo.NewClient(nil, nil, TOKEN_IPINFO)
	ip, err := client.GetIPAddr()
	if err == nil {
		fmt.Printf("ip: %v\n", ip)
		info, err := client.GetIPInfo(net.ParseIP(ip))
		if err != nil {
			fmt.Printf("err: %v\n", err)
		} else {
			fmt.Printf("info: %v\n", info)
			table.Append([]string{"IP", info.IP.String()})
			table.Append([]string{"City", info.City})
			table.Append([]string{"Region", info.Region})
			table.Append([]string{"Country", info.Country})
			table.Append([]string{"CountryName", info.CountryName})
			table.Append([]string{"Location", info.Location})
			table.Append([]string{"Org", info.Org})
			table.Append([]string{"Timezone", info.Timezone})
		}
	}

	table.Render()

	c.Print("           Host           ")
	fmt.Println("")

}
