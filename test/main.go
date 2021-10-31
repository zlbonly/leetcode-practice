package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	//testNums()
	//testNum4()
	//testNum5()

	nums := []int{0, 1, 2, 3, 4, 5, 5, 5, 5, 5, 5, 5, 6, 7, 8}
	fmt.Printf("%v", binarySearchFirst(nums, 9))
}

func binarySearchFirst(nums []int, target int) int {
	left, right, pivot := 0, len(nums)-1, 0
	for left < right {
		pivot = left + (right-left)/2
		if target > nums[pivot] {
			left = pivot + 1
		} else {
			right = pivot
		}
	}
	if nums[left] == target {
		return left
	}
	return -1
}

/**
升序数组热分查找
*/
func binarySearch(nums []int, target int) int {
	low, high, pivot := 0, len(nums)-1, 0
	for low <= high {
		pivot = low + (high-low)/2
		if target == nums[pivot] {
			return pivot
		} else if target > nums[pivot] {
			low = pivot + 1
		} else if target < nums[pivot] {
			high = pivot - 1
		}
	}
	return -1
}

func testNum5() {

	quit := make(chan struct{})

	go func() {
		fmt.Printf("%v", 22222222)
		quit <- struct{}{}
	}()
	//Label:
	for {
		select {
		case <-quit:
			return
			//break Label
		default:

		}
	}
}

func testNum4() {
	ch1, ch2, ch3 := make(chan bool), make(chan bool), make(chan bool)
	//var wg sync.WaitGroup
	//wg.Add(3)
	clo := make(chan bool)
	go func() {
		//Label:
		for i := 1; i <= 100; i += 3 {
			select {
			case <-ch1:
				fmt.Println("A -> ", i)
				if i < 100 {
					ch2 <- true
				} else {
					//close(ch1)
					//close(ch2)
					//close(ch3)
					//break Label
					//clo <- true
				}
			}
		}
		//wg.Done()
	}()
	go func() {
		//Label:
		for i := 2; i <= 100; i += 3 {
			select {
			case <-ch2:
				fmt.Println("B -> ", i)
				if i < 100 {
					ch3 <- true
				} else {
					//close(ch1)
					//close(ch2)
					//close(ch3)
					//break Label
					clo <- true
				}
			}
		}
		//wg.Done()
	}()
	go func() {
		//Label:
		for i := 3; i <= 100; i += 3 {
			select {
			case <-ch3:
				fmt.Println("C -> ", i)
				if i < 100 {
					ch1 <- true
				} else {
					//close(ch1)
					//close(ch2)
					//close(ch3)
					//break Label
					clo <- true
				}
			}
		}
		//wg.Done()
	}()
	ch1 <- true
Label:
	for {
		select {
		case <-clo:
			close(ch1)
			close(ch2)
			close(ch3)
			break Label
		default:

		}
	}
	//wg.Wait()
}

func testNums() {
	var wg sync.WaitGroup
	ch1 := make(chan bool, 1)
	ch2 := make(chan bool)
	ch3 := make(chan bool)

	wg.Add(3)
	go func() {
		defer func() {
			wg.Done()
		}()
		for i := 1; i <= 30; i += 3 {
			if _, ok := <-ch1; ok {
				fmt.Printf("协程1打印：%v\n", i)
				ch2 <- true
			}
		}
	}()

	go func() {
		defer func() {
			wg.Done()

		}()
		for i := 2; i <= 30; i += 3 {
			if _, ok := <-ch2; ok {
				fmt.Printf("协程2打印：%v\n", i)
				ch3 <- true
			}
		}
	}()

	go func() {
		defer func() {
			wg.Done()
			close(ch1)
			close(ch2)
			close(ch3)
		}()
		for i := 3; i <= 30; i += 3 {
			if _, ok := <-ch3; ok {
				fmt.Printf("协程3打印：%v\n", i)
				ch1 <- true
			}
		}
	}()
	ch1 <- true
	wg.Wait()
}

func findMinI(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		pivot := low + (high-low)/2
		if nums[pivot] < nums[high] {
			high = pivot
		} else {
			low = pivot + 1
		}
	}
	return nums[low]
}

/**

 */
func findMinUniqueII(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		pivot := low + (high-low)/2

		//  最小值一定在中间值nums[pivot]和最右侧值nums[pivot]中间，且不包括nums[pivot]，因此可以跳过。
		if nums[pivot] > nums[high] {
			low = pivot + 1
		} else if nums[pivot] < nums[high] {
			// 最小值一定在最左侧值nums[low]和中间值nums[pivot]之间，但是不确定nums[pivot]是否最小值，因此不能跳过。
			high = pivot
		} else {
			// 中间值nums[pivot]和最右侧值nums[high]相等，没法确定最小值位置，但是，nums[high]肯定有
			// nums[pivot]可以替换，因此可以忽略右端点。
			high--
		}
	}
	return nums[low]
}

func goroutineNums() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	goroutineTotals := 10
	ch := make(chan int, 3)
	wg := &sync.WaitGroup{}
	for i := 0; i < goroutineTotals; i++ {
		wg.Add(1)
		ch <- i
		go worker(ch, wg)
	}
	wg.Wait()
}

func worker(ch chan int, wg *sync.WaitGroup) {
	wg.Done()
	//time.Sleep(time.Second)
	if value, ok := <-ch; ok {
		fmt.Printf("执行时间:%v,worker%v\n", time.Now().UnixNano(), value)
	}
}

func controlChannelGoroutineNums() {
	count := 9
	limit := 3
	ch := make(chan bool, limit)
	fmt.Printf("%v hesh\n", runtime.NumCPU())
	runtime.GOMAXPROCS(4)
	var wg sync.WaitGroup
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(num int) {
			defer wg.Done()
			ch <- true
			fmt.Printf("%d 我在执行，time %d\n", num, time.Now().Unix())
			time.Sleep(2 * time.Second)
			<-ch
		}(i)
	}

	wg.Wait()
	close(ch)
}

func nums(num int) bool {
	if num < 0 {
		return false
	}

	total := 0

	cur := num

	for cur != 0 {

		total = total*10 + cur%10
		cur = cur / 10
	}

	fmt.Printf("%v ", total)

	return total == num
}
