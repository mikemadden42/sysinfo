package main

import (
	"fmt"
	"net"
	"os"
	"runtime"
)

func main() {
	cpuInfo()
	hostInfo()
	interfaceInfo()
	goInfo()
}

func cpuInfo() {
	fmt.Println("CPUs:", runtime.NumCPU())
}

func hostInfo() {
	hostname, _ := os.Hostname()
	fmt.Println("Hostname:", hostname)
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("OS arch:", runtime.GOARCH)
	fmt.Println("Temp dir:", os.TempDir())
}

func interfaceInfo() {
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

func goInfo() {
	fmt.Println("Go version:", runtime.Version())
}
