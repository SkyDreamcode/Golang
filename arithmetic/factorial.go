package arithmetic
import "fmt"

func Factorial(n uint64)(result uint64) {
	if(n > 0){
		result = n*Factorial(n - 1)
		return result
	}
	return 1
}

func ArithFactorial()(bool){
	var i int = 10
	fmt.Printf("%d factorial is %d\n", i, Factorial(uint64(i)))

	return true
}
