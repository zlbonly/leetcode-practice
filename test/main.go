package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func StudentRegister(name string, age int) *Student {
	s := new(Student)
	s.Name = name
	s.Age = age
	return s
}

func F() []int {
	a := make([]int, 0, 20)
	return a
}

func Slice() {
	s := make([]int, 10000, 100000)
	for index, _ := range s {
		s[index] = index
	}
}

func dynamic() {
	s := "Escape"
	fmt.Println(s)

	l := 20
	c := make([]int, 0, l) // 动态分配不定空间 逃逸
}
func main() {
	//StudentRegister("zlb",19) // 返回指针变量逃逸
	//F()
	//Slice()
	dynamic()
}

func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

/**

	一、逃逸场景
	1、指针逃逸
		Go可以返回局部变量指针，这其实是一个典型的变量逃逸案例
	通过go bulid -gcflags=-m 查看  发现 StudentRegister中 new(Student) 逃逸到堆上

	2、栈空间不足逃逸（空间开辟过大）
		实际上当栈空间不足以存放当前对象时或无法判断当前切片长度时会将对象分配到堆中。

	3、动态类型逃逸
		// 动态分配不定空间 逃逸
	例如 dynamic()

	4、闭包引用对象逃逸

	例如 Fibonacci
Fibonacci()函数中原本属于局部变量的a和b由于闭包的引用，不得不将二者放到堆上，以致产生逃逸。


二、逃逸分析的作用是什么呢？
逃逸分析的好处是为了减少gc的压力，不逃逸的对象分配在栈上，当函数返回时就回收了资源，不需要gc标记清除。

逃逸分析完后可以确定哪些变量可以分配在栈上，栈的分配比堆快，性能好(逃逸的局部变量会在堆上分配 ,而没有发生逃逸的则有编译器在栈上分配)。

同步消除，如果你定义的对象的方法上有同步锁，但在运行时，却只有一个线程在访问，此时逃逸分析后的机器码，会去掉同步锁运行。


三、逃逸总结：
1、栈上分配内存比在堆中分配内存有更高的效率

2、栈上分配的内存不需要GC处理

3、堆上分配的内存使用完毕会交给GC处理

4、逃逸分析目的是决定内分配地址是栈还是堆

5、逃逸分析在编译阶段完成

提问：函数传递指针真的比传值效率高吗？
我们知道传递指针可以减少底层值的拷贝，可以提高效率，但是如果拷贝的数据量小，由于指针传递会产生逃逸，可能会使用堆，也可能会增加GC的负担，所以传递指针不一定是高效的。





*/
