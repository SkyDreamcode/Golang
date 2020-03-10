package main //defaine package name

import "./myMath"//let the compiler know that program uses our own defined package  
import "fmt" //let compiler know that the program uses package fmt

func myPrint(){
	fmt.Println("start leran go program");
	fmt.Println(mathClass.Addmy(2,4));
	fmt.Println(mathClass.Submy(5,2));

}

func main(){
	myPrint();
	fmt.Println("end print");
	
}
