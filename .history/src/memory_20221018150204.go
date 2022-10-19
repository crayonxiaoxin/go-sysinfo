package src

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/gosuri/uilive"
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

	writer := uilive.New()
	table := KeyValueTable(writer)

	for {
		time.Sleep(time.Second)
		writer.Start()
		vms, err := mem.VirtualMemory()
		if err != nil {
			log.Fatal(err)
			return
		}
		table.AppendBulk([][]string{
			{"Total", SizeUnit(vms.Total)},
			{"Free", SizeUnit(vms.Free)},
			{"Available", SizeUnit(vms.Available)},
			{"Used", SizeUnit(vms.Used)},
			{"UsedPercent", strconv.FormatFloat(vms.UsedPercent, 'f', 4, 64) + " %"},
			{"Active", SizeUnit(vms.Active)},
			{"Inactive", SizeUnit(vms.Inactive)},
			{"Wired", SizeUnit(vms.Wired)},
			{"Buffers", SizeUnit(vms.Buffers)},
			{"Cached", SizeUnit(vms.Cached)},
			{"Shared", SizeUnit(vms.Shared)},
			{"SwapTotal", SizeUnit(vms.SwapTotal)},
			{"SwapCached", SizeUnit(vms.SwapCached)},
			{"SwapFree", SizeUnit(vms.SwapFree)},
		})
		table.Render()
		writer.Stop()
	}
}
