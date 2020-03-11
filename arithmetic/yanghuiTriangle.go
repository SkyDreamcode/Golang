package arithmetic
import "fmt"

func YhuiTriangle(n int)(bool){
	for i := 0; i< n; i++ {
		number := 1;
		for k := 0; k < n-i; k++ {
			fmt.Print(" ");
		}
		for j := 0; j <= i; j++ {
			fmt.Printf("%5d", number);
			number = number * (i - j)/(j + 1)
		}
		fmt.Println()
	}
	return true
}
