package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/fatih/color"
	"github.com/ipinfo/go/v2/ipinfo"
	"github.com/olekukonko/tablewriter"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

func main() {
	hostInfo()
	// cpuInfo()
	// memoryInfo()
	// diskInfo()
}

func hostInfo() {
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

	table.Render()

	c.Print("           Host           ")
	fmt.Println("")

	c3, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		defer c3.Close()
		fmt.Printf("c3: %v\n", c3.LocalAddr())
		u := c3.LocalAddr().(*net.UDPAddr)
		fmt.Printf("u: %v\n", u)
		fmt.Printf("c3: %v\n", c3.RemoteAddr())
	}

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

func memoryInfo() {
	vms, err := mem.VirtualMemory()
	if err != nil {
		log.Fatal(err)
		return
	}

	c := color.New(color.FgHiWhite)
	c = c.Add(color.BgBlue)
	c = c.Add(color.Bold)
	c.Print("           Memory           ")
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
	table.Append([]string{"Total", sizeUnit(vms.Total)})
	table.Append([]string{"Free", sizeUnit(vms.Free)})
	table.Append([]string{"Available", sizeUnit(vms.Available)})
	table.Append([]string{"Used", sizeUnit(vms.Used)})
	table.Append([]string{"UsedPercent", strconv.FormatFloat(vms.UsedPercent, 'f', 4, 64) + " %"})
	table.Append([]string{"Active", sizeUnit(vms.Active)})
	table.Append([]string{"Inactive", sizeUnit(vms.Inactive)})
	table.Append([]string{"Wired", sizeUnit(vms.Wired)})
	table.Append([]string{"Buffers", sizeUnit(vms.Buffers)})
	table.Append([]string{"Cached", sizeUnit(vms.Cached)})
	table.Append([]string{"Shared", sizeUnit(vms.Shared)})
	table.Append([]string{"SwapTotal", sizeUnit(vms.SwapTotal)})
	table.Append([]string{"SwapCached", sizeUnit(vms.SwapCached)})
	table.Append([]string{"SwapFree", sizeUnit(vms.SwapFree)})
	table.Render()

	c.Print("           Memory           ")
	fmt.Println("")
}

func diskInfo() {
	ps, err := disk.Partitions(false)
	if err != nil {
		log.Fatal(err)
		return
	}

	c := color.New(color.FgHiWhite)
	c = c.Add(color.BgBlue)
	c = c.Add(color.Bold)
	c.Print("           Disk           ")
	fmt.Println("")

	table := tablewriter.NewWriter(os.Stdout)
	header := []string{
		"Path",
		"Fstype",
		"Total",
		"Free",
		"Used",
		"UsedPercent(%)",
	}
	table.SetHeader(header)

	for _, partitionStat := range ps {
		us, err := disk.Usage(partitionStat.Mountpoint)
		if err != nil {
			continue
		}
		data := []string{
			us.Path,
			us.Fstype,
			sizeUnit(us.Total),
			sizeUnit(us.Free),
			sizeUnit(us.Used),
			strconv.FormatFloat(us.UsedPercent, 'f', 4, 64),
		}
		table.Append(data)
	}
	table.Render()

	c.Print("           Disk           ")
	fmt.Println("")
}

func sizeUnit(size uint64) string {
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
