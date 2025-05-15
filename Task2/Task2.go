package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
	"sync"
	"time"
)

var syncLock sync.Mutex

func main() {
	go func() {
		for i := 1; i <= 10; i += 2 {
			if i%2 != 0 {
				fmt.Printf("%d is odd\n", i)
			}
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for i := 2; i <= 10; i += 2 {
			if i%2 == 0 {
				fmt.Printf("%d is even\n", i)
			}
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(6 * time.Second)
	fmt.Printf("--------------------splitter--------------------")
	// 存放了所有任务的耗时数据
	spendMap := make(map[string]int64)
	// 定义了一个闭包切片
	var tasks []func(taskNum int)
	// 初始化闭包切片的值
	for i := 0; i < 10; i++ {
		task := func(taskNum int) {
			// 开始的毫秒数
			start := time.Now().UnixMilli()
			// 生成随机整数(100毫秒以内)，模拟任务执行的时间
			randomInt := rand.IntN(100)
			fmt.Printf("task #%d start\n", taskNum)
			time.Sleep(time.Duration(randomInt) * time.Millisecond)
			// 结束的毫秒数
			end := time.Now().UnixMilli()
			syncLock.Lock()
			// 计算用时，并存入map中
			spendMap["task"+strconv.Itoa(taskNum)] = int64(end - start)
			syncLock.Unlock()
		}
		tasks = append(tasks, task)
	}
	// 调用调度器调度任务并用协程并发执行
	schedule(tasks)
	time.Sleep(2 * time.Second)
	// 打印每个任务的执行耗时
	for k, v := range spendMap {
		fmt.Printf("%s spend %d millsecond \n", k, v)
	}
}

func schedule(tasks []func(taskNum int)) {
	for i := 0; i < len(tasks); i++ {
		go tasks[i](i + 1)
	}
}
