package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
)

func main() {
	cpuProfile := flag.String("cpuprofile", "", "write CPU profile to file")
	flag.Parse()

	if *cpuProfile != "" {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	v := []string{"hello", "world"}
	for x := range 10000 {
		for _, y := range v {
			fmt.Printf("%d-%s", x, y)
		}
	}
}
