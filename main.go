package main

import "fmt"
import "time" 

func main() {

	done := make (chan bool , 1) 

	go worker(done ) 
println(	<-done )


}

func worker (done chan bool ) { 
	fmt.Print("working... ")
	time.Sleep(time.Second) 
	fmt.Println("done"  )
	done <- true 

	

}

