package main 

import (
"fmt"

)

func main() {
	//fmt.Println("yogesh")
	ch := make(chan int)
go s(ch)

	for i := range ch {
	fmt.Println(i)
       }

}



func s(channel chan int) {
	for i := 0;i<=5;i++ {
	channel <- i
}
close(channel)

}