package varPrint
import "fmt"
import "unsafe"

func calculateArea(){
	const LENGTH int = 100
	const HIGH = 23
	
	var area int
	area = LENGTH*HIGH
	fmt.Println("The area is :");
	fmt.Printf("%d\n", area);
}

//likeness enum of c language
const(
	female = 1
	male = 2
	unknow =0
)

const(
	gca = "golang"
	gcb = len(gca)
	gcc = unsafe.Sizeof(gca)
)

const(
	gc1 = iota   //0
	gc2          //1
	gc3          //2
	gc4 = "hellow"
	gc5          //4
	gc6 = "word"
	gc7          //6
	gc8 = 100
        gc9 = iota    //8
        gc10          // 9
)

const(
	i = 1<<iota //1<<0 = 1
	j = 2<<iota //2<<1 = 4
	k           //2<<2 = 8
	l           //2<<3 = 16
)


func ConstPrint(){
	//explicitly define
	const b string = "golang hello"
	//implicitly define
	const a = "golang word"

	const c,d,e = "hhhh", 2 ,50
	fmt.Println("a=",a, "b=", b, "c=", c, "d=", d, "e=", e)
	
	calculateArea();
	
	fmt.Println("female:", female, " male:", male, " unknow:", unknow);
	fmt.Println("gca:", gca, " gcb", gcb, " gcc", gcc);
	fmt.Println("gc1:",gc1, " gc2", gc2, " gc3", gc3, " gc4", gc4, " gc5", gc5, " gc6", gc6, " gc7", gc7, " gc8", gc8, " gc9", gc9, " gc10", gc10);
	fmt.Printf("i=%d, j=%d, k=%d, l=%d\n", i,j,k,l);
}
