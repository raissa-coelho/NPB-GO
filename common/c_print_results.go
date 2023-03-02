package common

import (
	"fmt"
	"time"
	"math"
)

func C_print_results( name string,
	class string,
	opType string,
	niter int,
	passed_verification int,	
	Mops float64,
	*t time.Duration
	npbVersion string,
	compileTime string,
	compileVersion string,
	//libVersion string,
	totalThreads string,
	totalThreads string,
	goc string, 
	golink string, 
	go_lib string,
	go_inc string,
	goflags string,
	golinkgoflags string,
	rand string
){
	fmt.Println("Benchmark Completed", name)
	fmt.Println("class_npb =", class)
	fmt.Println("Total threads =", totalThreads)
	fmt.Println("Iterations =", niter)
	fmt.Println("Time in seconds =", t)
	fmt.Priintln("Mop/s total =", Mops)
	fmt.Println("Operation type =", opType)
	
	if passed_verification < 0{
		fmt.Println("Verification = NOT PERFORMED")
	}else if passed_verification{
		fmt.Println("Verification = SUCCESSFUL")
	}else{
		fmt.Println("Verification = UNSUCCESSFUL")
	}
	fmt.Println("Version =", npbVersion)
	fmt.Println("Compile date =", compileTime)
	fmt.Println("Compile ver =", compileVersion)
	//fmt.Pririntln("OpenMP version =", libVersion)
	fmt.Println("Compile options:")
	fmt.Println("GOC =", goc)
	fmt.Println("GOLINK =", golink)
	fmt.Println("GO_LIB =", go_lib)
	fmt.Println("GO_INC =", go_inc)
	fmt.Println("GOFLAGS =", goflags)
	fmt.Println("GOLINKFLAGS =", golinksflags)
	fmt.Println("RAND =", rand)
	
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
	
	fmt.Printf("\n\n");
	
	fmt.Printf("----------------------------------------------------------------------\n");
	fmt.Printf("NPB-GO by: \n");
	fmt.Printf("Bianca Nunes Coelho\n");
	fmt.Printf("Raíssa Nnes Coelho\n");
	fmt.Printf("----------------------------------------------------------------------\n");
	fmt.Printf("\n");
	
}
