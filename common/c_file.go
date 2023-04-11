package common

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"
)

func Guarda(t *time.Duration) {
	var aux string
	if runtime.GOOS == "windows" {
		aux = "..\\bin\\"
	} else {
		aux = "../bin/"
	}
	f, err := os.OpenFile(aux+"guarda.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	f.WriteString(fmt.Sprint(*t) + "\n")

	defer f.Close()
}

func C_file(bench string, M int, gc, sx, sy float64, NQ int, q []float64, rng string, nit int, verified bool, Mops float64, t *time.Duration) {
	var aux string
	if runtime.GOOS == "windows" {
		aux = "..\\bin\\"
	} else {
		aux = "../bin/"
	}

	var n string
	if M == 24 {
		n = "S"
	} else if M == 25 {
		n = "W"
	} else if M == 28 {
		n = "A"
	} else if M == 30 {
		n = "B"
	} else if M == 32 {
		n = "C"
	} else if M == 36 {
		n = "D"
	} else if M == 40 {
		n = "E"
	}

	f, err := os.Create(aux + bench + "_" + n + ".txt")
	if err != nil {
		log.Fatal(err)
	}

	f.WriteString("Benchmark :" + bench + "\n")
	f.WriteString("N = " + strconv.Itoa(M) + "\n")
	f.WriteString("No. Gaussian Pairs = " + fmt.Sprint(gc) + "\n")
	f.WriteString("Sums = " + fmt.Sprint(sx) + "  " + fmt.Sprint(sy) + "\n")
	f.WriteString("Counts: \n")
	for i := 0; i < NQ-1; i++ {
		f.WriteString(strconv.Itoa(i) + " - " + fmt.Sprint(q[i]) + "\n")
	}
	f.WriteString("Benchmark Completed!\n")
	f.WriteString("Class = " + n + "\n")
	f.WriteString("Total threads = " + strconv.Itoa(runtime.NumCPU()) + "\n")
	f.WriteString("Iterations = " + strconv.Itoa(nit) + "\n")
	f.WriteString("Time = " + fmt.Sprint(*t) + "\n")
	f.WriteString("Mop/s total = " + fmt.Sprint(Mops) + "\n")
	f.WriteString("Operation type = " + rng + "\n")

	if verified {
		f.WriteString("Verification = SUCCESSFUL\n")
	} else {
		f.WriteString("Verification = UNSUCCESSFUL\n")
	}

	f.WriteString("Compiler Version = " + fmt.Sprint(runtime.Version()))

	defer f.Close()
}

func C_file_IS(passed bool, bench, class string, TOTAL_KEYS, MAX_ITERATIONS int, Mops float64, t *time.Duration) {
	var aux string
	if runtime.GOOS == "windows" {
		aux = "..\\bin\\"
	} else {
		aux = "../bin/"
	}

	f, err := os.Create(aux + bench + "_" + class + ".txt")
	if err != nil {
		log.Fatal(err)
	}

	f.WriteString("NAS Parallel Benchmark Parallel GO version - IS Benchmark\n")
	f.WriteString("Class: " + class + "\n")
	f.WriteString("Size: " + strconv.Itoa(TOTAL_KEYS) + "\n")
	f.WriteString("Iterations: " + strconv.Itoa(MAX_ITERATIONS) + "\n")
	f.WriteString("Number of available goroutines: " + strconv.Itoa(runtime.NumCPU()) + "\n")
	f.WriteString("\n\n     iteration    \n")
	for i := 0; i < MAX_ITERATIONS; i++ {
		f.WriteString("\n\n     " + strconv.Itoa(MAX_ITERATIONS))
	}
	f.WriteString("\n\n")
	f.WriteString("Benchmark Completed!\n")
	f.WriteString("Class = " + class + "\n")
	f.WriteString("Size = " + strconv.Itoa(TOTAL_KEYS) + "\n")
	f.WriteString("Iterations = " + strconv.Itoa(MAX_ITERATIONS) + "\n")
	f.WriteString("Time in seconds = " + fmt.Sprint(*t) + "\n")
	f.WriteString("Total goroutines = " + strconv.Itoa(runtime.NumCPU()) + "\n")
	f.WriteString("Mop/s Total = " + fmt.Sprint(Mops) + "\n")
	Mops_pt := Mops / float64(runtime.NumCPU())
	f.WriteString("Mop/s/goroutine  = " + fmt.Sprint(Mops_pt) + "\n")
	f.WriteString("Operation type = keys ranked\n")
	if passed {
		f.WriteString("Verification = SUCCESSFUL\n")
	} else {
		f.WriteString("Verification = UNSUCCESSFUL\n")
	}
	f.WriteString("Compiler Version = " + fmt.Sprint(runtime.Version()))
}
