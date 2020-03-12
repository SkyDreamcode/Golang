package interfacePro
import "fmt"

type student interface{
	studentName() string
	studentAge() int
}

type highStudent struct {
	name string
	age int
}

type collegeStudent struct {
	name string
	age int

}
func (hStudent highStudent) studentName()string{
	//fmt.Println("this is high student")
	hStudent.name = "apple1"
	return hStudent.name
}

func (hStudent highStudent) studentAge()int {
	hStudent.age = 66
	return hStudent.age
}

func (cStudent collegeStudent) studentName() string{
	//fmt.Println("this is college student")
	cStudent.name = "grape"
	return cStudent.name
}

func (cStudent collegeStudent) studentAge() int {
	cStudent.age = 80
	return cStudent.age
}


func InterfaceTest() bool{
	var stu student
	
	stu = new(highStudent)
	fmt.Printf("high name:%s, age:%d\n", stu.studentName(), stu.studentAge())

	stu = new(collegeStudent)
	fmt.Printf("college name:%s, age:%d\n", stu.studentName(), stu.studentAge())
	
	return true
}
