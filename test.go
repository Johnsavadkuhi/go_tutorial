package main 

import "fmt"

func main() {

	go hello() 
	fmt.Println("main function ")

}

func hello () { 
	fmt.Println("hello world from mohammad ") 
}