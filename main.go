package main

import "log"

func main()  {
	for n:=range sq(gen(5,4,7,9,15,11,12,17)){
		log.Println(n)
	}

}

func gen(nums ...int) <-chan int {
	out:=make(chan int,len(nums))
	go func() {
		for _,n:=range nums{
			out<-n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int{
	out:=make(chan int)
	go func() {
		for n:=range in{
			out<-n*n
		}
		close(out)
	}()
	return out
}

