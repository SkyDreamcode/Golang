package structTest
import "fmt"

type book struct{
	title string
	author string
	subject string
	book_id int
}

func booktest(book1 book,book2 book)(book, book){
	book1.title = "golang"
	book1.author = "abcd"
	book1.subject = "sub1"
	book1.book_id = 1

	book2.title = "cpp"
	book2.author = "author2"
	book2.subject = "sub2"
	book2.book_id = 2

	return book1,book2	
}

func booktest1(book1 *book,book2 *book){
	book1.title = "golang"
	book1.author = "abcd"
	book1.subject = "sub1"
	book1.book_id = 1

	book2.title = "cpp"
	book2.author = "author2"
	book2.subject = "sub2"
	book2.book_id = 2

	//return book1,book2	
}

func StructTest(arg int)(int){
	//create a new struct
	fmt.Println(book{"golang1", "kawayi", "init", 3487});
	//key : value
	fmt.Println(book{title:"golang2",author:"abcd", subject:"init", book_id:45677});
	//
	fmt.Println(book{title:"golang3",author:"kkkk"});

	var book1 book
	var book2 book
	
	//book1,book2 = booktest(book1, book2)
	booktest1(&book1, &book2)
	
	fmt.Printf("book1: title:%s author:%s subject:%s book_id:%d\n", book1.title, book1.author, book1.subject, book1.book_id);
	fmt.Printf("book2: title:%s author:%s subject:%s book_id:%d\n", book2.title, book2.author, book2.subject, book2.book_id);

	return arg
}
