package ep

import (
	"fmt"
	"os"
	"math"
	"time"
)

//Defining constants
const MK = int(16)
const MM = M - MK
const NN = 1 << MM
const NK = 1 << MK
const NQ = 10
const EPSILON = 1.0e-8
const A = 1220703125.0
const S = 271828183.0
const NK_PLUS = ((2*NK)+1)

func EP(){	
	var x = [NK_PLUS]float64{}
	var q = [NQ]float64{}
	var Mops, sx, sy ,an, gc, t1 float64
	var np, nit int
	dum := [3]float64{1.0,1.0,1.0}
	
	
	//Time variable
	var t time.Duration
	//Variables for verification of success
	var verified bool
	var sx_err float64
	var sy_err float64
	var sx_verify_value float64
	var sy_verify_value float64	

	fmt.Println("NAS Parallel Benchmark Parallel GO version - EP Benchmark")
	fmt.Println("Number of random numbers generated", math.Pow(2.0,M+1))
	
	verified = false
	
	np = NN

	r.Vrancl(0, &dum[0], dum[1], dum[2])
	dum[0] = r.Randlc(&dum[1], dum[2])
	for i := 0; i < NK_PLUS; i++{
		x[i] = -1.0e99
	}
	Mops = math.Log(math.Sqrt(math.Max(1.0,1.0)))

	t1 = A
	r.Vranlc(0,&t1,A,x)
	t1 = A
	
	for i := 0; i < MK-1; i++{
		r.Randlc(&t1,t1)
	} 
	
	an = t1
	t = S
	gc = 0.0
	sx = 0.0
	sy = 0.0
	
	for i := 0; i <= NQ-1; i++{
		q[i] = 0.0
	}
	
	//programação paralela aqui
		

	for i := 0; i < NQ-1; i ++{
		gc = gc + q[i]
	}
	
	//Verification of the values.
	nit = 0
	verified = true
	
	if M == 24 {
		sx_verify_value = -3.247834652034740e+3
		sy_verify_value = -6.958407078382297e+3
	}else if  M == 25 {
		sx_verify_value = -2.863319731645753e+3
		sy_verify_value = -6.320053679109499e+3
	}else if M == 28 {
		sx_verify_value = -4.295875165629892e+3
		sy_verify_value = -1.580732573678431e+4
	}else if M == 30 {
		sx_verify_value =  4.033815542441498e+4
		sy_verify_value = -2.660669192809235e+4
	}else if M == 32 {
		sx_verify_value =  4.764367927995374e+4
		sy_verify_value = -8.084072988043731e+4
	}else if M == 36 {
		sx_verify_value =  1.982481200946593e+5
		sy_verify_value = -1.020596636361769e+5
	}else if M == 40 {
		sx_verify_value = -5.319717441530e+05
		sy_verify_value = -3.688834557731e+05
	}else {
		verified = false
	}
	
	if verified == true {
		sx_err = math.Abs((sx - sx_verify_value) / sx_verify_value)
		sy_err = math.Abs((sy - sy_verify_value) / sy_verify_value)
		verified = (sx_err <= EPSILON) && (sy_err <= EPSILON)
	}
	Mops = math.Pow(2.0, float64(M+1))/t.Seconds/1000000.0	
		
	//Print of the results of the benchmark.
	 fmt.Println("EP Benchmark Results:")
	 fmt.Printf("CPU Time = %v\n", tm)	
	 fmt.Printf("N = %v\n", M)
	 fmt.Printf("No. Gaussian Pairs = %v\n", gc)
	 fmt.Printf("Sums = %v %v\n", sx,sy)
	 fmt.Printf("Counts: \n")
	 for i := 0; i < NQ-1; i++{
	 	fmt.Printf("%v - %v\n", i, q[i])
	 }
		 

}
