package main

import "strings"


func main(){

	sentences := []string{"Hello Reza" , "Say Love Me" , "Take care of yourself", "this is a book from library God"}

	NumWords := MostWordFound(sentences)
	println(NumWords)
}

func MostWordFound(sentences []string) int {
    
	max := 0 
	for _ , item := range sentences {
		length := len(strings.Split(item , " "))
		if length > max {
			max = length
		}

		

	}
	return max

}