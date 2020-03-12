package concurrencePro
import (
	"fmt"
	"time"
)

func setData(s string, c chan int){
	for i := 0; i < 30; i++{
		time.Sleep(100*time.Millisecond)
		c <- i
		fmt.Println(s)
	}
}

func getData(c chan int){
	for i := 0; i < 20; i++{
		time.Sleep(100*time.Millisecond)
		fmt.Println("getData from chan:", <-c)
	}
}

func conTest(){
	c := make(chan int, 10)
	go setData("apple", c)
	getData(c)
}

func fibonacci(n int, c chan int){
	x, y := 0, 1
	for i := 0; i < n; i++{
		c <- x
		x, y = y, x + y
	}
	close(c)
}

func fibonacciTest(){
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	
	for i := range c{
		fmt.Println(i)
	}
}

func ConcurrenceTest()bool{
	conTest()	
	fibonacciTest()	
	return true
}
