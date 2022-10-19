package src

import (
	"fmt"
	"log"
	"strconv"

	"github.com/fatih/color"
	"github.com/shirou/gopsutil/v3/mem"
)

func MemoryInfo() {
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

	table := KeyValueTable()

	table.Append([]string{"Total", SizeUnit(vms.Total)})
	table.Append([]string{"Free", SizeUnit(vms.Free)})
	table.Append([]string{"Available", SizeUnit(vms.Available)})
	table.Append([]string{"Used", SizeUnit(vms.Used)})
	table.Append([]string{"UsedPercent", strconv.FormatFloat(vms.UsedPercent, 'f', 4, 64) + " %"})
	table.Append([]string{"Active", SizeUnit(vms.Active)})
	table.Append([]string{"Inactive", SizeUnit(vms.Inactive)})
	table.Append([]string{"Wired", SizeUnit(vms.Wired)})
	table.Append([]string{"Buffers", SizeUnit(vms.Buffers)})
	table.Append([]string{"Cached", SizeUnit(vms.Cached)})
	table.Append([]string{"Shared", SizeUnit(vms.Shared)})
	table.Append([]string{"SwapTotal", SizeUnit(vms.SwapTotal)})
	table.Append([]string{"SwapCached", SizeUnit(vms.SwapCached)})
	table.Append([]string{"SwapFree", SizeUnit(vms.SwapFree)})
	table.Render()

	c.Print("           Memory           ")
	fmt.Println("")
}
