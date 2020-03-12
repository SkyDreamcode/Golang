package primeNumber
import "fmt"

func PrimeNumber(){
	var a ,b int
	for a = 2; a <= 100; a++ {
		for b = 2; b <= (a/b); b++ {
			if 0 == a%b{
				break
			}
		}
		if b > (a/b){
			fmt.Printf("%d \t is prime number\n", a)
		}
	}
}

func PrimeNumber1(x int) (bool){
	var a, b int;
	var c bool;
	c = true;
	for a = 2; a <= x; a++{
		for b = 2; b < a; b++{
			if(0 == a%b){
				c = false;
			}
		}
		if(false != c){
			fmt.Printf("func1:%d\t is prime number\n", a);
		}else{
			c = true;
		}
	}
	return c
}
