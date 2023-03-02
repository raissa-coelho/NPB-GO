package main

import(
	"fmt"
	"os"
	"time"
	"log"
)

var benchmark_types [10]string = {SP, BT, LU, MG, FT, IS, EP, CG, UA, DC}

func main(){
	
	var typeBench int
	var class string
	//var class_old string
	
	//Verify number of arguments passed in the command line
	if len(os.Args) != 4 {
		fmt.Println("Usage: make benchmark_name CLASS=benchmark_type ")
	}
	args := os.Args
	
	//Verify if arguments are right
	get_info(args,&typeBench,&class)
	if class != "U"{
		fmt.Println("setparams: For benchmark ",args[2],"class = ",class)
		check_info(typeBench,class)
	}

	//read_info(typeBench, &class_old)
	write_info(typeBench, class)
	
	//if class != class_old {
	//	fmt.Printf("Writing...")
	//	write_info(typeBench,class)
	//}else{
	//	fmt.Println("setparams: File unmodified.")
	//}


}

//Verify if benchmark_name is ok
func get_info(args string,*typeBench int, *classp string){

	*classp = args[3]
	
	if (args[2] == "SP" || args[2] == "sp") {
	   *typeBench = 0
	}else if (args[2] == "BT" || args[2] == "bt") {
	   *typeBench = 1
	}else if (args[2] == "LU" || args[2] == "lu") {
	   *typeBench = 2
	}else if (args[2] == "MG" || args[2] == "mg") {
	   *typeBench = 3
	}else if (args[2] == "FT" || args[2] == "ft") {
	   *typeBench = 4
	}else if (args[2] == "IS" || args[2] == "is") {
	   *typeBench = 5
	}else if (args[2] == "EP" || args[2] == "ep") {
	   *typeBench = 6
	}else if (args[2] == "CG" || args[2] == "cg") {
	   *typeBench = 7
	}else if (args[2] == "UA" || args[2] == "ua") {
	   *typeBench = 8
	}else if (args[2] == "DC" || args[2] == "dc") {
	   *typeBench = 9
	}else{
		fmt.Println("setparams: Error: unknown benchmark type: " + args[1])
		os.Exit(1)
	}
}

//Verify if benchmark_type is ok
func check_info(typeB int, class string){
	if (class != "S" && class != "W" && class != "A" && class != "B" && class != "C" && class != "D" && class != "E" && class != "F"){
		fmt.Println("setparams: Unknown benchmark class ",class)
		fmt.Println("setparams: Allowed classes are S, W, A, B, C, D, E and F.")
		os.Exit(1)
	}
}

//Read file of a previous benchmark
func read_info(typeB int, classOld string){
	//to be done at a later date
}

//Write settings
func write_info(typeB int, class string){
	file, err := os.Create("npbparams.txt")
	if err != nil {
		log.Fatal(err)
	}
	
	_ , err2 := file.WriteString("//THIS FILE CAN NOT BE CHANGED!\n")
	if err2 != nil {
		log.Fatal(err2)
	}
	
	switch typeB {
	
	//EP
	case 6: writeEP(file, class)	   
	
	default: fmt.Println("setparams: Error. Unknown benchmark type.")
		 f.Close()
		 os.Exit(1)
	}		
}

//EP benchmark information
func writeEP(f *os.File, class string ){
	var M int
	if class == "S"{
	   M = 24
	}else if class == "w" {
		M = 25
	}else if class == "A" {
		M = 28
	}else if class == "B" {
		M = 30 
	}else if class == "C" {
		M = 32
	}else if class == "D" {
		M = 36
	}else if class == "E" {
		M = 40
	}else {
		fmt.Println("setparams: Internal error: invalid class type")
		os.Exit(1)
	}
	defer f.CLose()
	_,err1 := f.WriteString(class)
	if err1 != nil{
		log.Fatal(err1)
	}
	
	_,err2 := f.WriteString(string(M))
	if err2 != nil{
		log.Fatal(err1)
	}	
}
