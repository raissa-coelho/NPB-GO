package main

import(
	ep "NPB-GO/EP"
	"fmt"
	"strconv"
	"os"
	"runtime"
	"log"
)

func init(){
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main(){
	
	var typeBench int
	var class string
	
	//Verify number of arguments passed in the command line
	if len(os.Args) != 3 {
		fmt.Println("Usage: make benchmark_name CLASS=benchmark_type ")
	}
	args := os.Args
	
	//Verify if arguments are right
	get_info(args,&typeBench,&class)
	if class != "U"{
		fmt.Println("setparams: For benchmark",args[1],"class= ",class)
		check_info(typeBench,class)
	}

	write_info(typeBench, class)

}

//Verify if benchmark_name is ok
func get_info(args[] string, typeBench *int, classp *string){

	*classp = args[2]
	
	if (args[1] == "SP" || args[1] == "sp") {
	   *typeBench = 0
	}else if (args[1] == "BT" || args[1] == "bt") {
	   *typeBench = 1
	}else if (args[1] == "LU" || args[1] == "lu") {
	   *typeBench = 2
	}else if (args[1] == "MG" || args[1] == "mg") {
	   *typeBench = 3
	}else if (args[1] == "FT" || args[1] == "ft") {
	   *typeBench = 4
	}else if (args[1] == "IS" || args[1] == "is") {
	   *typeBench = 5
	}else if (args[1] == "EP" || args[1] == "ep") {
	   *typeBench = 6
	}else if (args[1] == "CG" || args[1] == "cg") {
	   *typeBench = 7
	}else if (args[1] == "UA" || args[1] == "ua") {
	   *typeBench = 8
	}else if (args[1] == "DC" || args[1] == "dc") {
	   *typeBench = 9
	}else{
		fmt.Println("setparams: Error: unknown benchmark type: ", args[1])
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
		 file.Close()
		 os.Exit(1)
	}		
}

//EP benchmark information
func writeEP(f *os.File, class string){
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
	defer f.Close()
	_,err1 := f.WriteString(class)
	if err1 != nil{
		log.Fatal(err1)
	}
	_, err3 := f.WriteString("\n")
	if err3 != nil{
		log.Fatal(err3)
	}
	_,err2 := f.WriteString(strconv.Itoa(M))
	if err2 != nil{
		log.Fatal(err1)
	}	
	
	ep.Ep(M)
}
