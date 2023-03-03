package common

import (
	"fmt"
	"runtime"
	"time"
)

func C_print_results(class string, opType string, niter int, passed_verification bool, Mops float64, t *time.Duration, totalThreads string) {
	fmt.Printf("Benchmark Completed!\n")
	fmt.Printf("Class NPB =%v\n", class)
	fmt.Printf("Total threads =%v\n", totalThreads)
	fmt.Printf("Iterations =%v\n", niter)
	fmt.Printf("Time in seconds =%v\n", *t)
	fmt.Printf("Mops total =%v\n", Mops)
	fmt.Printf("Operation type =%v\n", opType)

	if passed_verification {
		fmt.Println("Verification = SUCCESSFUL")
	} else {
		fmt.Println("Verification = UNSUCCESSFUL")
	}

	fmt.Println("Compiler Version =", runtime.Version())

	/*
	 * fmt.Printf(" Please send the results of this run to:\n\n");
	 * fmt.Printf(" NPB Development Team\n");
	 * fmt.Printf(" Internet: npb@nas.nasa.gov\n \n");
	 * fmt.Printf(" If email is not available, send this to:\n\n");
	 * fmt.Printf(" MS T27A-1\n");
	 * fmt.Printf(" NASA Ames Research Center\n");
	 * fmt.Printf(" Moffett Field, CA  94035-1000\n\n");
	 * fmt.Printf(" Fax: 650-604-3957\n\n");
	 */

	fmt.Printf("\n\n")

	fmt.Printf("----------------------------------------------------------------------\n")
	fmt.Printf("NPB-GO by: \n")
	fmt.Printf("Bianca Nunes Coelho\n")
	fmt.Printf("Raíssa Nunes Coelho\n")
	fmt.Printf("----------------------------------------------------------------------\n")
	fmt.Printf("\n")

}
