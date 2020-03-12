package main //defaine package name

//import "./myMath"//let the compiler know that program uses our own defined package  
//import "./printTest"
//import "./varprint"
//import "./primeNumber"
//import "./arithmetic"
//import "./structtest"
//import "./slice"
import( 
	"fmt" //let compiler know that the program uses package fmt
	//"./range"
	//"./arithmetic"
	//"./interface"
	"./concurrence"
)

/*
func myPrint(){
	fmt.Println("start leran go program");
	fmt.Println(mathClass.Addmy(2,4));
	fmt.Println(mathClass.Submy(5,2));

}
*/

func main(){
	//myPrint();
	//fmt.Println("myPrint end print");
	//strPrint.StrPrint();
	//varPrint.VarPrint();
	//varPrint.GlobleVarTest();
	//varPrint.ConstPrint();
	//primeNumber.PrimeNumber();
	//var c bool
	//c = primeNumber.PrimeNumber1(50);
	//fmt.Println("c=", c);
	
	//var b bool
	//b = arithmetic.YhuiTriangle(12);
	//fmt.Println("b = ", b);

	var a bool
	//a = structTest.StructTest(12)
	//a = slicePro.SliceTest()
	//a = rangePro.RangeTest()
	//a = arithmetic.ArithFactorial()
	//a = arithmetic.ArithFibonacci()
	//a = interfacePro.InterfaceTest()
	a = concurrencePro.ConcurrenceTest()
	fmt.Println("a = ", a)
}
