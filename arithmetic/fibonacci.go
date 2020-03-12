package arithmetic
import "fmt"

func ArithFibonacci() bool{
	for i := 0; i < 10; i++{
		fmt.Printf("%d\t", fibonacci(i))
	}
	return true
}

func fibonacci(i int) int{
	if(i < 2){
		return i
	}
	return fibonacci(i - 2) + fibonacci(i - 1)
}
