package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Println("Temp dir:", os.TempDir())
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Go version:", runtime.Version())
}
