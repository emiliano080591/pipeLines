package main

import "log"

// TODO: Build a Pipeline
// generator() -> square() -> print
// set up the pipeline

// run the last stage of pipeline
// receive the values from square stage
// print each one, until channel is closed.
func main()  {
	for n:=range square(generator(5,4,7,9,15,11,12,17)){
		log.Println(n)
	}

}


