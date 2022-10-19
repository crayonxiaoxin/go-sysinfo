package src

import (
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/fatih/color"
	"github.com/ipinfo/go/v2/ipinfo"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
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

	infoStats, err := cpu.Info()
	if err == nil {
		count := len(infoStats)
		for i := 0; i < count; i++ {
			cpu := infoStats[i]
			index := ""
			if count > 1 {
				index = strconv.FormatInt(int64(cpu.CPU), 10) + " "
			}
			table.AppendBulk([][]string{

				{"Cores" + index, strconv.FormatInt(int64(cpu.Cores), 10)},
				{"CacheSize" + index, strconv.FormatInt(int64(cpu.CacheSize), 10)},
				{"ModelName" + index, cpu.ModelName},
				{"Mhz" + index, strconv.FormatFloat(cpu.Mhz, 'f', 2, 64)},
			})
			if count > 1 {
				table.Append([]string{"", ""})
			}
		}
	}

	vms, err := mem.VirtualMemory()
	if err == nil {
		table.AppendBulk([][]string{
			{"Memory", SizeUnit(vms.Total)},
		})
	}

	us, err := disk.Usage("/")
	if err == nil {
		table.AppendBulk([][]string{
			{"Disk", SizeUnit(us.Total)},
		})
	}

	client := ipinfo.NewClient(nil, nil, TOKEN_IPINFO)
	ip, err := client.GetIPAddr()
	if err == nil {
		info, err := client.GetIPInfo(net.ParseIP(ip))
		if err != nil {
			fmt.Printf("err: %v\n", err)
		} else {
			table.AppendBulk([][]string{
				{"", ""},
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
