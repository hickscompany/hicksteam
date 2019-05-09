package main

import "fmt"

type books struct{		//定义一个结构体类型

	title string
	author string
	subject string
	book_id int
}
//阶乘函数
func Factorial(n int)(result int){

	
	if (n>0){

		fmt.Println(n)
		result = n*Factorial(n-1)

		fmt.Printf("%x\n",&result)
		return result
	}
	return 1
}

//斐波那契数列

func fibonacci(n int) (result int){

	if n<2{

		return n


	}
		return fibonacci(n-2)+fibonacci(n-1)
}


//定义接口
//接口传参，及返回结果
type phone interface{

	call(param int)  string

	takephone()
}
//定义类型
type samsung struct{

}

//实现接口方法 
func (samsungphone samsung) call(param int) string{

	fmt.Println("i am sumsung, i can call you!",param)
	return "demo"
}

func (samsungphone samsung) takephone(){
	fmt.Println("i can take a photo for you")
}
func main(){


	var n[10]int
	var i int

	for i := 0; i < 10; i++ {
	 	
	 	n[i] =i+100
	 	fmt.Printf("Element[%d]=%d\n",i,n[i])

	} 
	fmt.Println(i,n)

	a:=20

	ip :=&a

	fmt.Printf("a 的变量地址是:%x\n",&a)

	fmt.Printf("ip 变量储存指针的地址是:%x\n",ip)

	 //*用于获取指针指向的内容
	fmt.Println("*ip 变量的值是",*ip)


	//创建一个新的结构体
	fmt.Println(books{"go语言","hicks","www.iii",34})

	// 也可以使用 key => value 格式 忽略字段0或空
	fmt.Println(books{author:"golang"})



	var book1 books //声明 book1 的类型为books
	var book2 books //声明 book2 的类型为books

	/**
	 * book1描述
	 */
	book1.title="golang"
	book1.author="googel"
	book1.subject="go"
	book1.book_id=1	 

	/* book 2 描述 */
	book2.title = "Python 教程"
	book2.author = "www.runoob.com"
	book2.subject = "Python 语言教程"
	book2.book_id = 6495700

	/* 打印 book1 信息 */
	fmt.Printf( "book 1 title : %s\n", book1.title)
	fmt.Printf( "book 1 author : %s\n", book1.author)
	fmt.Printf( "book 1 subject : %s\n", book1.subject)
	fmt.Printf( "book 1 book_id : %d\n", book1.book_id)

	/* 打印 book2 信息 */
	fmt.Printf( "book 2 title : %s\n", book2.title)
	fmt.Printf( "book 2 author : %s\n", book2.author)
	fmt.Printf( "book 2 subject : %s\n", book2.subject)
	fmt.Printf( "book 2 book_id : %d\n", book2.book_id)




	//递归


	var dg=15

	fmt.Printf("%d 的阶乘是 %d\n",dg,Factorial(dg))


	for i := 0; i < 10; i++ {

		fmt.Printf("%d\t",fibonacci(i))
		
	}


	var phone phone
	phone =new (samsung)
	phone.takephone()
	r:=phone.call(500)
	fmt.Println(r)


	//测试闭包函数原理代码 闭包函数变量从外部获取导致close i =4
	var fs=[4]func(){}


		for i := 0; i < 4; i++ {

		fs[i] = func(){fmt.Println("close i",i)}
		
	}

	for _,f:=range fs{

		f()
	}



}			