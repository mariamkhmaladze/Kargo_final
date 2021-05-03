package main

import (
	"fmt"
	"os"
	//"strconv"
)
var NumberToWord=map[byte]string{
               '0':"Zero",
               '1':"One",
               '2':"Two",
               '3':"Three",
               '4':"Four",
               '5':"Five",
               '6':"Six",
               '7':"Seven",
               '8':"Eight",
               '9':"Nine",
               


}

func main() {
        problem1 := os.Args[1]
        problem2 := os.Args[2]
        problem3 := os.Args[3]

        var intArray []string
	intArray=append(intArray, problem1, problem2, problem3)
	
	str:=""
	for i:=0; i<len(intArray); i++{
	          s:=intArray[i]
	          for j:=0; j<len(s); j++{
	               str+=NumberToWord[s[j]]
	
	          }
	          str+=","
	
	
	}
    
	fmt.Printf("%q\n", str[:(len(str))-1])
}

