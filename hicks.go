package main

import (
	"fmt"
	"strconv" //int 和 string 做类型转换需要使用到的包
	"sort" //排序包
)

type(  //自定义变量类型别名
	文本 string
)

const(

	A = 1  //在定义常量组时，如果不提供初始值，则表示将使用上行的表达式

	B = 0

	C,D = iota,2  //iota为枚举转换
)

func main() {

	var a bool //布尔类型只有true或false。不能使用数字代替

	var b[1]int //[]中的是改变量的容量 数组

	var c 文本

	c = "变态"

	d :=3.333 //最简单声明赋值变量

	var e,f,g =1,2,3 //局部变量只能并行简写。或简写成 e,f,g :=1,2,3

	h := strconv.Itoa(e) //类型转换需要显示的表达变量的类型 h := strconv.Itoa(e) 作为类型转换。直接输出为数字。因为计算机本身所以的都会认为是数字

	// h = "sb" //为h重新赋值 下面的赋值会覆盖掉上方已有变量的值


	//流程控制语句
	if i :=1;i<=3{
		
		fmt.Println("if is work")
	}

	//for循环语句
	for i :=1; i<=10; i++{

		d++

		fmt.Println(d)
	}
	//选择语句switch
	switch a:=11;{

		case a>0:
			fmt.Println(a)

			fallthrough //fallthrough关键字使用可以继续执行下一个语句块
		case a>=10:
			fmt.Println(a+1)
		default:
			fmt.Println("over")

	}

//跳转语句

	
	fmt.Println("fk")
	fmt.Println("34")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e,f,g,h,A,B,C,D)
LABLE1:
	
	for{	
		for i:=1;i<=10;i++{

			if i>3{

				break LABLE1 //如果是goto应该把LABLE1放在for之下 因为goto是跳到指定区域继续执行，break是跳出该区域 continue也是跳到外层继续执行但是外层for必须是有限循环
			}
		}
	}
	//冒泡排序
	k :=[...] int{5,9,2,6,3,1}
	fmt.Println(k)

	num :=len(k)

	for i := 0; i < num; i++ {

		for j := i+1; j < num; j++ {

			if k[i]>k[j] {

				temp:=k[i]
				k[i]=k[j]
				k[j]=temp
				
			}
			
		}
		
	}
	fmt.Println(k)
	
	//slice 切片
	array :=[]int{3,4,7,8,9,10} //非数组。指向底层数组

	// l :=[]int{1,2,3,4,5,6}
 
	l1 :=array[3:5]  //指向数组

	l2 :=array[2:6]

	fmt.Println("the l2 is ",l2)

	l2 =append(l2,2,3,4,5,6,7,8,11,11,11,1,1)

	l1[0]=2

	fmt.Println("the slice",l1,cap(l2),array) //其中一个slice指向小标值改变会引发全部改变 ,如果超出底层数组容量将不再指向原array底层数组，而是重新分配内存地址到新的底层数组

	s1 :=make([]int,3,6) //初始化  3为下标长度，6为slice容量

	fmt.Printf("%p\n",s1)

	s1 =append(s1,1,2,3,4,5) //追加数组数量超过cap容量会导致重新分配地址

	fmt.Printf("%p,%v\n",s1,s1)

	//map 关键字
	m := make(map[int]string) //make关键字进行map的初始化

	m[1] = "ok" 

	delete(m,1) //delete 删除map中的对应键值

	m1 :=map[int]string{1:"a",2:"b",3:"c",4:"d"}

	sm1 := make([]int,len(m1)) //进行slice的初始化方便后面对map的值键值进行存放

	mi := 0

	for k,_:=range m1{

		sm1[mi] = k

		mi++ 
	}

	m2 := make(map[string]int)

	for k,v:=range m1{

		m2[v] = k
							//对两组map的键值对调
	}
	sort.Ints(sm1)
	fmt.Println(m,m1,m2,sm1)


	//调用函数部分
	sa :=[]int{2,9,0}

	Afunc(sa)

	sa1 :=func(){
		fmt.Println("work")
	}
	sa1() //匿名函数

	fmt.Println(sa)


	for i := 0; i < 4; i++ {

		defer func(){

			fmt.Println(i)
		}()
		
		
	}

	Bfunc()	

	Cfunc()

	var fs = [4]func(){}

	for i1 := 0; i1 < 4; i1++ {
		
		defer fmt.Println("defer i=",i1)
		defer func(){fmt.Println("defer closure i=",i1)}()

		fs[i1] = func(){fmt.Println("closure i= ",i1)}
	}	


	for _,f:=range fs{

		f()
	}

}
// 函数

func Afunc(s []int){

	//如果调用函数里的直接参数是slice的话是会改变调用前的参数的。因为slice是地址拷贝
	s[0] =3
	s[1] =5
	s[2] =8
	fmt.Println(s)
} 

func Bfunc(){

	defer func(){

		if err:= recover(); err!= nil{

			fmt.Println("recover in B")
		}
	}()

	panic("panic Bfunc")
}

func Cfunc(){

	fmt.Println("Print Cfunc")
}


//change

















