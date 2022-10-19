package src

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/fatih/color"
	"github.com/ipinfo/go/v2/ipinfo"
	"github.com/olekukonko/tablewriter"
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

	table := tablewriter.NewWriter(os.Stdout)
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

	table.Render()

	c.Print("           Host           ")
	fmt.Println("")

	client := ipinfo.NewClient(nil, nil, "5a1445e6b70829")
	ip, err := client.GetIPAddr()
	if err == nil {
		fmt.Printf("ip: %v\n", ip)
		c2, err := client.GetIPInfo(net.ParseIP(ip))
		if err != nil {
			fmt.Printf("err: %v\n", err)
		} else {
			fmt.Printf("c2: %v\n", c2)
		}
	}
}
