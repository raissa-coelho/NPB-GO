//------------------------------------------------------------------------------
//The original NPB 3.4.1 version was written in Fortran and belongs to: 
//	http://www.nas.nasa.gov/Software/NPB/
//Authors of the Fortran code:
//	P. O. Frederickson
//	D. H. Bailey
//	A. C. Woo
//	H. Jin
//------------------------------------------------------------------------------
//The serial C++ version is a translation of the original NPB 3.4.1
//Serial C++ version: https://github.com/GMAP/NPB-CPP/tree/master/NPB-SER
//Authors of the C++ code: 
//	Dalvan Griebler <dalvangriebler@gmail.com>
//	Gabriell Araujo <hexenoften@gmail.com>
// 	Júnior Löff <loffjh@gmail.com>
//------------------------------------------------------------------------------

//GO Language Version - EP Benchmark - TECVIII
// Bianca Nunes Cooelho
// Raíssa Nunes coelho

package EP

import (
	r "NPB-GO/common"
	"fmt"
	"os"
	"math"
	"time"
	"runtime"
)

//Defining constants
const MK = int(16)
const NK = 1 << MK
const NQ = 10
const EPSILON = 1.0e-8
const A = 1220703125.0
const S = 271828183.0
const NK_PLUS = ((2*NK)+1)


func Ep(class string, M int){
	var MM = M - MK
	var NN = 1 << MM	
	
	var x = [NK_PLUS]float64{}
	var q = [NQ]float64{}
	var Mops, sx, sy ,an, gc, t1 float64
	var np, nit int
	dum := [3]float64{1.0,1.0,1.0}
	var k_offset int
	
	//Time variable
	var t time.Duration
	//Variables for verification of success
	var verified bool
	var sx_err float64
	var sy_err float64
	var sx_verify_value float64
	var sy_verify_value float64	
	
	fmt.Println("NAS Parallel Benchmark Parallel GO version - EP Benchmark")
	fmt.Println("Number of random numbers generated", math.Pow(2.0,float64(M+1)))
	
	verified = false
	
	np = NN

	r.Vrandlc(0, &dum[0], dum[1], dum[2])
	dum[0] = r.Randlc(&dum[1], dum[2])
	for i := 0; i < NK_PLUS; i++{
		x[i] = -1.0e99
	}
	Mops = math.Log(math.Sqrt(math.Max(1.0,1.0)))

	t1 = A
	r.Vrandlc(0,&t1,A,x[:])
	
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
	
	//GO Channels - buffered
	qR := make([]chan float64, np)
	for i:= 0; i<len(qR);i++{
		qR[i] = make(chan float64)
	}
	sxR := make(chan float64, np)
	syR := make(chan float64, np)
	
	//Begining of parallel programing
	k_offset = -1
	start := time.Now()
	for k := 1; k <= np; k++{
		go func(k int){
			//Temporary varible
			var t1,t2,t3,t4,x1,x2 float64
			var kk, ik, l int
			var qq = [NQ]float64{}
			var x = [NK_PLUS]float64{}
			kk = k_offset + k
			t1 = S
			t2 = an
			
			for i:=1; i <=100; i++{
				ik = kk/2
				if ((2*ik) != kk){
					t3 = r.Randlc(&t1,t2)
				}
				if (ik == 0){
					break
				}
				t3 = r.Randlc(&t2,t2)
				kk = ik
			}
			r.Vrandlc((2*NK), &t1, A, x)
			
			for i := 0; i< NK; i++{
				x1 = 2.0 * x[2*i] - 1.0
				x2 = 2.0 * x[2*i+1] - 1.0
				t1 = math.Pow(x1,2) + math.Pow(x2,2)
				if (t1 <= 1.0){
					t2 = math.Sqrt((-2.0) * math.Log(t1) / t1)
					t3 = (x1 * t2)
					t4 = (x2 * t2)
					l = int(math.Max(math.Abs(t3), math.Abs(t4)))
					qq[l] += 1.0
					sx = sx + t3
					sy = sy + t4
				}
			}
			syR <- sy
			sxR <- sx
			qR <- qq 
		}(k)
		
		for i := 0; i<= np; i++{
			sx <- sxR
			sy <- syR
			for j := 0; j < NQ-1; j++{
				q[j] <- qR[j]
			}
		}
	
	} 
	//End of parrallel programing
	stop := time.Now()
	*t = stop.Sub(start)
	
	close(qR)
	close(sxR)
	close(syR)
	
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
	Mops = math.Pow(2.0, float64(M+1))/ (*t)/1000000.0	
		
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
	
	
	r.C_print_results( "EP",class,"Random Numbers Generated",nit,verified,Mops,&t,runtime.NumCPU())
		 
}
