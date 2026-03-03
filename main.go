package main

import (
	"1brc/runner"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	var r runner.Runner
	var runnerInp string
	var inpFile string
	var outFile string

	args := os.Args
	if len(args) < 3 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}
	runnerInp = args[1]
	inpFile = args[2]
	outFile = args[3]

	switch runnerInp {
	case "v1":
		r = runner.V1{}
	default:
		fmt.Printf("incorrect input %v", runnerInp)
		os.Exit(1)
	}

	_, ok := os.LookupEnv("DEBUG")
	if ok {
		cpuProf, err := os.Create("cpu.prof")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		defer cpuProf.Close()

		pprof.StartCPUProfile(cpuProf)
	}

	now := time.Now()
	res := r.Run(inpFile)
	dur := time.Since(now)
	fmt.Printf("Took %.2f seconds\nWriting output\n", dur.Seconds())

	if ok {
		pprof.StopCPUProfile()
	}

	f, err := os.Create(outFile)
	if err != nil {
		log.Panicln(err.Error())
	}
	defer f.Close()

	for k, v := range res {
		v[2] = v[1] / v[0]
		f.WriteString(fmt.Sprintf("%v,%v,%v\n", k, v[1], v[2]))
	}
}
