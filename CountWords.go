package main

import (
	"fmt"
	"strings"
)
 
func CountWords(text string)int{

	if text==""{  
		return 0
	}
	input:=strings.Split(text, " ")

	count:=0
	for _, word:= range input{

		if word != ""{
			count++
		}
	}
	return count
}
func main(){
	fmt.Println(CountWords("the boy  is a good boy in go"))
}
