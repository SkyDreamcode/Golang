package rangePro
import "fmt"


func RangeTest()(bool){
	var x int
	var y string
        nums := []int{1,2, 3, 4, 5, 6, 7, 8, 9, 0}

        sum := 0
        for _, num := range nums {
                sum += num
        }
        fmt.Println("sum=", sum)

        for i, num := range nums {
                fmt.Printf("nums[%d] = %d\n", i, nums[i:(i+1)])
		num += 1
        }
        //range can map
        kvs := map[string]string{"a":"apple", "b":"banana"}
        for k,v := range kvs{
                fmt.Printf("%s -> %s\n", k, v)
        }

	mt := make(map[int]string)
	mt[0] = "France"
	mt[1] = "japan"
	mt[2] = "korea"
	mt[3] = "russia"
	mt[4] = "american"
	mt[5] = "england"
 	mt[6] = "iran"
	
	for x, y= range mt{
		fmt.Printf("mt:%d -> %s\n", x, y)
	}

	delete(mt, 1)
	fmt.Println("delete 1")
	for x, y= range mt{
		fmt.Printf("mt:%d -> %s\n", x, y)
	}

        //range can enum unicode string,args1:the index of string, arg2:the string(unicode) itself
        for i, c := range "golang"{
                fmt.Println("enum unicode :", i, c)
        }

        return true
}
