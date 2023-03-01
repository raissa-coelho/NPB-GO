package common

import (
	"fmt"
	"os"
	"math"
)

func C_print_results( name string,
	class,opType string
	n1,n2,n3,niter,passed_verification int	
	Mops,t float64
	npbVersion,compileTime,compileVersion,libVersion,totalThreads string
	totalThreads, goc, golink, go_lib, go_inc, goflags, golinkgoflags,rand string
){
	fmt.Println("Benchmark Completed", name)
	fmt.Println("class_npb =", class)
	
	if (name[0] == "I" && name[1] == "S"){
		if n3 == 0{
			nn := n1
			if n2 != 0{
				nn *= n2
			}
			fmt.Println("Size =", nn)
		}else{
			fmt.Println("Size =", n1,n2,n3)
		}	
	}else{
		size string
		j int
		if (n2==0 && n3==0){
			if (name[0] == "E" && name[1]=="P"){
				fmt.Println(size,math.Pow(2.0,n1))
				j = 14
				if size[j] == "."{
					size[j] = " "
					j--
				}
				fmt.Println("Size =", size)
			}else{
				fmt.Println("Size =", n1)
			}
		
		}else{
			fmt.Println("Size =", n1,n2,n3)
		}
	}
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
	fmt.Printf("NPB-GO is developed by: \n");
	fmt.Printf("Bianca NUnes Coelho\n");
	fmt.Printf("Raíssa Nnes Coelho\n");
	fmt.Printf("\n");
	fmt.Printf("In case of questions or problems, please send an e-mail to us:\n");	
	fmt.Printf("bncoelho; rncoelho@inf.ufpel.edu.br\n");
	fmt.Printf("----------------------------------------------------------------------\n");
	fmt.Printf("\n");
	
}
