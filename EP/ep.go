package ep

import (
	"fmt"
	"os"
)

func EP(){







	//Verification of the values.
	verified := true
	
	if M == 24 {
		sx_verify_value = -3.247834652034740e+3;
		sy_verify_value = -6.958407078382297e+3;
	}else if  M == 25 {
		sx_verify_value = -2.863319731645753e+3;
		sy_verify_value = -6.320053679109499e+3;
	}else if M == 28 {
		sx_verify_value = -4.295875165629892e+3;
		sy_verify_value = -1.580732573678431e+4;
	}else if M == 30 {
		sx_verify_value =  4.033815542441498e+4;
		sy_verify_value = -2.660669192809235e+4;
	}else if M == 32 {
		sx_verify_value =  4.764367927995374e+4;
		sy_verify_value = -8.084072988043731e+4;
	}else if M == 36 {
		sx_verify_value =  1.982481200946593e+5;
		sy_verify_value = -1.020596636361769e+5;
	}else if M == 40 {
		sx_verify_value = -5.319717441530e+05;
		sy_verify_value = -3.688834557731e+05;
	}else {
		verified := false
	}
	
	if verified == true {
		sx_err =
	}
	
	//Print of the results of the benchmark.
	 fmt.Println("EP Benchmark Results:")
	 fmt.Printf("CPU Time = %v\n", tm)	
	 fmt.Printf("N = %v\n", M)
	 fmt.Printf("No. Gaussian Pairs = %v\n", gc)
	 fmt.Printf("Sums = %v %v\n", sx,sy)
	 fmt.Printf("Counts: \n")

}
