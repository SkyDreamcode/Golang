package slicePro
import "fmt"

func printSlice(x []int, i int){
	fmt.Printf("index: %d  len:%d cap:%d s:%v\n", i, len(x), cap(x), x)
}

func appendTest(s []int){
	printSlice(append(s, 0), 0) //allow additional empty slice
	printSlice(append(s, 1), 1) //add an element to the s slice
	printSlice(append(s, 2,3, 4), 2)//add multi-element to the slice

	//var num []int
//	num := make([]int, len(s), (cap(s))*2)//create the slice that it is double cap than before the s
//	copy(num, s)
//	printSlice(num, 4)
}

func SliceTest()(bool){
	var number = make([]int, 4, 8)
	printSlice(number, 0)
	if(number == nil){
		fmt.Printf("the slice:number is empty\n")
	}
	
	var arr = [...]int{11,12, 13, 14, 15, 16, 17, 18} 

	s1 := []int {1,2,3}
	fmt.Println("s1 :", s1)
	s2 := arr[:]
	fmt.Println("s2:", s2[:], "--", s2[3:], "--", s2[3:5])
	s3 := arr[3: 6]
	s4 := arr[3:]
	s5 := arr[:6]
	s6 := s3[4: 5]
	s7 := make([]int, 5, 7)
	
	printSlice(s1, 1)
	printSlice(s2, 2)
	printSlice(s3, 3)
	printSlice(s4, 4)
	printSlice(s5, 5)
	printSlice(s6, 6)
	printSlice(s7, 7)

	appendTest(s2);


	return true
}

