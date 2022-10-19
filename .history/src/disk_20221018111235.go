package src

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/shirou/gopsutil/v3/disk"
)

func DiskInfo() {
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
			SizeUnit(us.Total),
			SizeUnit(us.Free),
			SizeUnit(us.Used),
			strconv.FormatFloat(us.UsedPercent, 'f', 4, 64),
		}
		table.Append(data)
	}
	table.Render()

	c.Print("           Disk           ")
	fmt.Println("")
}
