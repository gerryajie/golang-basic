package main

import (
	"fmt"
	"pertama/calculation"
)
func main(){
	fmt.Println("halo,belajar golang")
	
	result:=calculation.Add(9,9)
	kali:=calculation.Multy(3,5)

	fmt.Println(result)
	fmt.Println(kali)
	}