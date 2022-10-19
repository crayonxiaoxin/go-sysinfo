package src

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/shirou/gopsutil/v3/cpu"
)

func CpuInfo() {
	infoStats, err := cpu.Info()
	if err != nil {
		log.Fatal(err)
		return
	}

	c := color.New(color.FgHiWhite)
	c = c.Add(color.BgBlue)
	c = c.Add(color.Bold)
	c.Print("           CPU           ")
	fmt.Println("")

	count := len(infoStats)
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
	for i := 0; i < count; i++ {
		cpu := infoStats[i]
		index := ""
		if count > 1 {
			index = strconv.FormatInt(int64(cpu.CPU), 10) + " "
		}
		table.Append([]string{"Cores" + index, strconv.FormatInt(int64(cpu.Cores), 10)})
		table.Append([]string{"ModelName" + index, cpu.ModelName})
		table.Append([]string{"Mhz" + index, strconv.FormatFloat(cpu.Mhz, 'f', 2, 64)})
		if count > 1 {
			table.Append([]string{"", ""})
		}
	}

	table.Append([]string{"", ""})
	table.Render()

	c.Print("           CPU           ")
	fmt.Println("")
}
