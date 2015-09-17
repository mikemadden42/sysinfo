package main

import (
	"fmt"
	"net"
	"os"
	"runtime"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Go version:", runtime.Version())
	fmt.Println("OS arch:", runtime.GOARCH)
	fmt.Println("OS type:", runtime.GOOS)
	fmt.Println("Temp dir:", os.TempDir())

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
