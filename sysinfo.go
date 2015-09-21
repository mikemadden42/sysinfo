package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"net"
	"os"
	"runtime"
)

func main() {
	cpu_info()
	mem_info()
	host_info()
	interface_info()
	go_info()
}

func cpu_info() {
	fmt.Println("CPUs:", runtime.NumCPU())
	c, _ := cpu.CPUInfo()
	for index, each := range c {
		fmt.Printf("CPU %d speed: %f Mhz, Cache: %d\n", index, each.Mhz, each.CacheSize)
	}
}

func mem_info() {
	v, _ := mem.VirtualMemory()
	fmt.Printf("Total Memory: %v, Free: %v, UsedPercent: %f%%\n", v.Total, v.Free, v.UsedPercent)

	s, _ := mem.SwapMemory()
	fmt.Printf("Total Swap: %v, Free: %v, UsedPercent: %f%%\n", s.Total, s.Free, s.UsedPercent)
}

func host_info() {
	h, _ := host.HostInfo()
	fmt.Printf("Hostname: %s\n", h.Hostname)
	fmt.Println("OS arch:", runtime.GOARCH)
	fmt.Printf("OS: %s\n", h.OS)
	fmt.Printf("Platform Family: %s\n", h.PlatformFamily)
	fmt.Printf("Platform: %s %s\n", h.Platform, h.PlatformVersion)
	fmt.Println("Temp dir:", os.TempDir())
}

func interface_info() {
	interfaces, _ := net.Interfaces()
	for _, inter := range interfaces {
		if inter.Name != "lo" {
			fmt.Println(inter.Name, inter.HardwareAddr)
		}
		if addrs, err := inter.Addrs(); err == nil {
			for _, addr := range addrs {
				fmt.Println(inter.Name, "->", addr)
			}
		}
	}
}

func go_info() {
	fmt.Println("Go version:", runtime.Version())
}
