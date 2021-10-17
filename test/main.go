package main

import (
	"fmt"
	"sync"
	"time"
	"unsafe"
)

/**
	channel使用场景
	1、信号通知 （一个任务完成了，通知另一个任务执行 。例如：经常会有这样的场景，当信息收集完成，通知下游开始计算数据）
	2、超时处理
	3、生产消费模型
	4、数据传递
	5、控制并发数
	6、自定义互斥锁（通过设置一个缓冲区为1的通道，如果成功的往通道里发送数据，说明拿到锁，否则锁被别人拿走了，等待他人解锁）


	channel产生的panic
	1、关闭1个nil值的channel会引发panic
	2、关闭一个已关闭的channel会引发panic
	3、向一个已关闭的channel发送数据

	1、向已关闭的channel写入数据会触发panic
	2、从关闭的channel读取
		1、有缓存通道 。依旧可以读出关闭前写入的值，如果没有值时，会返回该类型的零值
		2、无缓存通道。读出对应类型的零值
·
	select可同时监控多路channel，并处理最先发生的channel

	为什么需要协程池？
	1、无限开启协程，导致标准输出被过多并发操作，从而产生panic。
	2、无限开启协程，导致程序占用过多内存，存在内存不足而崩溃的风险。
	3、协程过多会造成gc的压力。

	1、Ants - 高性能低损耗的 Goroutine 池
		Ants 对于任务的执行原理比较直观，通过一个工作池的形式维护 goroutine 集合。
		当向工作池提交任务时，从池中取出 worker 来执行。如果已经存在可用的 goroutine 了，
		那么直接开始执行，如果没有，则需要判断是否已经达到容量上限。如果还没有超过，
		那就意味着可用的 worker 比容量更少，此时启动新的 worker 来执行。
		而如果容量已经用完，就依据是否为阻塞模式，来马上返回，或是阻塞等待。

		当任务执行完毕，对应的 worker 就会得到释放，重新回到池中，等待下一个任务的调度，实现 goroutine 的复用。


*/

func main() {
	//testNotify()
	testOdd()
}

func testOdd() {
	var wg sync.WaitGroup
	ch := make(chan struct{})

	wg.Add(2)

	// 1、打印奇数
	go func() {
		defer wg.Done()
		for i := 1; i < 101; i++ {
			ch <- struct{}{}
			if i%2 == 1 {
				fmt.Printf("%d ", i)
			}
		}
	}()

	// 2、打印偶数
	go func() {
		defer wg.Done()
		for i := 1; i < 101; i++ {
			<-ch
			if i%2 == 0 {
				fmt.Printf("%d ", i)
			}
		}
	}()
	wg.Wait()
}

func testNotify() {
	ch := make(chan struct{})
	go func() {
		collectMsg(ch)
	}()
	<-ch

	calculateMsg()
}

func calculateMsg() {
	fmt.Println("开始进行数据分析~")
}

func collectMsg(isOver chan struct{}) {
	fmt.Println("开始采集数据")
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("完成采集")
	isOver <- struct{}{}
}

func testChannel1() {
	fmt.Println("Begin doing something")
	c := make(chan bool)

	go func() {
		defer close(c)
		fmt.Println("doing something")
	}()

	<-c

	fmt.Println("Done ! ")
}
func worker(start chan bool, index int) {
	<-start
	fmt.Println("This is worker:", index)
}

func testChanncel2() {
	start := make(chan bool)
	for i := 1; i <= 100; i++ {
		go worker(start, i)
	}
	close(start)

	select {}
}

func producer(c chan int, max int) {
	defer close(c)
	for i := 0; i < max; i++ {

		c <- i
	}
}

func consumer(c chan int) {
	var v int
	ok := true
	for ok {
		if v, ok = <-c; ok {
			fmt.Println(v)
		}
	}
}

var quit chan int = make(chan int)

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d", i)
	}
	quit <- 0
}

/**
系统的线程会抢占式地输出

*/

/***
1、缓存为1的通道和无缓存通道的区别
答：1、无缓存通道channel必须在接受方与发送方同时准备好时，通道才能正常传递数据，否则双方只有一方在线都会阻塞
	无缓存通道如果将将接收方与发送方放在一个程序里，会死锁
	2、有缓存channel,当缓冲区满时，发送数据会阻塞，当缓存区为空时，接受数据会阻塞。发送方和接收方不需要同时做好准备。
*/
func main11() {

	//done2 :=make(chan bool,1)
	//fmt.Print(done1,done2)

	/*var(
		name string
		age int
		desc string
	)
	//fmt.Scan(&name,&age,&desc)

	fmt.Scanln(&name,&age,&desc)

	fmt.Printf("%v%v%v",name,age,desc)*/
	/*go loop()
	go loop()

	for i:=0;i<2;i++{
		<-quit
	}
	time.Sleep(10e6)

	defer func() {
		if r := recover();r != nil{
			fmt.Printf("%v",111)
		}
	}()*/
	/*ch :=make(chan int,10)
	go producer(ch,30)
	go consumer(ch)

	time.Sleep(10e6)*/
	//testChannel1()
	//testChanncel2()

	/*f1 :="yyyy-MM-dd"
	f2 :="M/d/yyyy"
	f3 := "yyyy/M/d"
	s :=""
	fmt.Scan(&s)
	indexArr :=make([]int,0)
	for i:=0;i<len(s);i++{
		fmt.Printf("%v",string(s[i]))
		if string(s[i]) == "'"{
			indexArr = append(indexArr,i)
		}
	}
	json.Marshal()

	s1 := s[indexArr[0]+1:indexArr[1]]
	s2 := s[indexArr[2]+1:indexArr[3]]
	sArr1 := strings.Split(s1,"/")
	if s2 == f1 {
		s2 = strings.Replace(s2,"yyyy",sArr1[0],1)
		s2 = strings.Replace(s2,"MM",sArr1[1],1)
		s2 = strings.Replace(s2,"dd",sArr1[2],1)

	}else if  s2 == f2 {
		s2 = strings.Replace(s2,"M",sArr1[1],1)
		s2 = strings.Replace(s2,"d",sArr1[2],1)
		s2 = strings.Replace(s2,"yyyy",sArr1[0],1)
	}else if s2 == f3 {
		s2 = strings.Replace(s2,"yyyy",sArr1[0],1)
		s2 = strings.Replace(s2,"M",sArr1[1],1)
		s2 = strings.Replace(s2,"d",sArr1[2],1)
	}*/
}

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
}

type UserData struct {
	Name string
}

func GetUserInfo(userInfo UserData) *UserData {
	return &userInfo
}

/***

优化建议

func main() {
	var info UserData
	info.Name = "WilburXu"
	_ = GetUserInfo(&info)
}

func GetUserInfo(userInfo *UserData) *UserData {
	return userInfo
}


*/

type demo1 struct {
	a int8
	b int16
	c int32
}

type demo2 struct {
	a int8
	c int32
	b int16
}

func main1() {
	/*t := time.After(time.Second * 3)
	fmt.Printf("t type=%T\n", t)
	//阻塞3秒
	fmt.Println("t=", <-t)*/
	/*	a :=[]int{1,2,3,4,5,6,7,8}
		fmt.Printf("%v",ZlbSerarch(len(a), func(i int) bool {
			return a[i] >= 2
		}))*/
	fmt.Println(unsafe.Alignof(demo1{}))
	fmt.Println(unsafe.Alignof(demo2{}))

	fmt.Printf("demo sizeof =%v,demo2 sizeof = %v", unsafe.Sizeof(demo1{}), unsafe.Sizeof(demo2{}))
}

func ZlbSerarch(n int, f func(int) bool) int {

	i, j := 0, n

	for i < j {
		h := int(uint(i+j) >> 1)
		if !f(h) {
			i = h + 1
		} else {
			j = h
		}
	}
	return i
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
