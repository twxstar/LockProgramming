package cas

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var threadNum int64
var curNum int64
var threadWorkGroup sync.WaitGroup

func atmoicOperation() {
	var loopParam int64
	for {
		//等待所有协程都完成
		if threadNum == 10 {
			break
		}
	}

	for {
		//测试对threadNum的并发操作
		loopParam = atomic.LoadInt64(&curNum)
		//atomic.CompareAndSwapInt64具有三个参数，第一个是变量的地址，第二个是变量当前值，第三个是要修改变量为多少，该函数如果发现传递的old值等于当前变量的值，则使用第三个变量替换变量的值并返回true，否则返回false。
		result := atomic.CompareAndSwapInt64(&curNum, loopParam, loopParam+1)
		//如果成功，则跳出循环
		if result {
			//fmt.Println(loopParam, " Success Try to CAS: ", result)
			break
		} else {
			fmt.Println(loopParam, " Error Try to CAS: ", result)
		}
	}

	threadWorkGroup.Done()
}

func SingleTestComapareAndSwap() {
	var oldValue int32 = 1
	var loopParam int32 = 2
	result := atomic.CompareAndSwapInt32(&oldValue, loopParam, loopParam+1)

	println(result)
}

func DoWork() {
	curNum = 0
	threadNum = 0
	for i := 0; i < 10; i++ {
		go atmoicOperation()
		threadWorkGroup.Add(1)
		threadNum += 1
	}

	threadWorkGroup.Wait()

}
