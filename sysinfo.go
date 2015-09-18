package main

import (
	"fmt"
	"net"
	"os"
	"runtime"
)

func main() {
	cpu_info()
	os_info()
	interface_info()
	go_info()
}

func cpu_info() {
	fmt.Println("CPUs:", runtime.NumCPU())
}

func os_info() {
	fmt.Println("OS arch:", runtime.GOARCH)
	fmt.Println("OS type:", runtime.GOOS)
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
