package varPrint
import "fmt"

var x, y int
var (
	ga int
	gb bool
)

var gc,gd int = 1,2
var ge,gf = 3444, "Golang hello"

//only local variables
//g,h := 3,5


func VarPrint(){
	var a string = "golang"
	fmt.Println(a);

	var intb int
	fmt.Println(intb);

	var c bool
	fmt.Println(c);


	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("the var test:i=%v f=%v b=%v s=%q\n", i, f, b, s);
	fmt.Println("the var test:i=%v f=%v b=%v s=%q\n", i, f, b, s);


	var d = true //omit the type of the variable , default bool type
	fmt.Println("d = ", d)
	
	var intval int
	//intval := 1//compile error:no new variables on left side of :=
	intval,intval1 := 4,5
	intval2 := 6
	fmt.Println("intval=", intval, "intval1=", intval1, "intval2=", intval2);
	strtest := "golang program"
	fmt.Println("strtest=", strtest); 


	lg,lh := true, 234
	fmt.Println("print variables:");
	fmt.Println(x, y, ga, gb, gc, gd, ge, gf, lg,lh);
	fmt.Println("print variables:");
	lg,lh = false, 666
	fmt.Println(x, y, ga, gb, gc, gd, ge, gf, lg,lh);
	fmt.Println("function test:");
	lj,lk := funtest();
	fmt.Printf("lj=%v, lk=%s\n", lj, lk);
	_,lk = funtest();
	fmt.Printf("lk=%s\n", lk);
}



func funtest()(int ,string){
	a, b := 23, "function test"
	return a,b
}
