package main

import "log"

func main()  {
	in:=generator(2,3,4,5,6,7,8)

	ch1:=square(in)
	ch2:=square(in)

	for n:=range merge(ch1,ch2){
		log.Println(n)
	}
}
