package main

import (
	"crayonxiaoxin/sysinfo/src"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "sysinfo"}
	var cmdHost = &cobra.Command{
		Use:     "host",
		Aliases: []string{"h"},
		Example: "sysinfo -host",
		Run: func(cmd *cobra.Command, args []string) {
			src.HostInfo()
		},
	}
	var cmdCpu = &cobra.Command{
		Use:     "cpu",
		Aliases: []string{"c"},
		Example: "sysinfo -cpu",
		Run: func(cmd *cobra.Command, args []string) {
			src.CpuInfo()
		},
	}
	var cmdDisk = &cobra.Command{
		Use:     "disk",
		Aliases: []string{"d"},
		Example: "sysinfo -disk",
		Run: func(cmd *cobra.Command, args []string) {
			src.DiskInfo()
		},
	}
	var cmdMem = &cobra.Command{
		Use:     "memory",
		Aliases: []string{"m", "mem"},
		Example: "sysinfo memory || sysinfo menm",
		Run: func(cmd *cobra.Command, args []string) {
			src.DiskInfo()
		},
	}
	rootCmd.AddCommand(cmdHost, cmdCpu, cmdDisk, cmdMem)
	rootCmd.Execute()
}
