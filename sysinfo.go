package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Go version:", runtime.Version())
	fmt.Println("OS arch:", runtime.GOARCH)
	fmt.Println("OS type:", runtime.GOOS)
	fmt.Println("Temp dir:", os.TempDir())
}
