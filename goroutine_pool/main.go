package main

import (
	"fmt"
	"math/rand"
)

type Job struct {
	Id      int
	RandNum int
}

type Result struct {
	Job *Job
	sum int
}

func main() {
	// job管道
	jobChan := make(chan *Job, 128)
	// 结果管道
	resultChan := make(chan *Result, 128)
	// 创建工作池
	createPool(64, jobChan, resultChan)
	// 开启打印协程
	go func(resultChan chan *Result) {
		// 遍历结果管道并打印
		for result := range resultChan {
			fmt.Printf("job %v random:%v get result %d\n", result.Job.Id, result.Job.RandNum, result.sum)
		}
	}(resultChan)
	var id int
	// 创建循环job，输入到管道
	for {
		id++
		// 生成随机数
		r_num := rand.Int()
		job := &Job{Id: id, RandNum: r_num}
		jobChan <- job
	}

}

func createPool(num int, jobChan chan *Job, resultChan chan *Result) {
	// 根据开协程个数，去跑运行
	for i := 0; i < num; i++ {
		go func(JobChan chan *Job, resultChan chan *Result) {
			// 执行运算
			// 遍历job管道所有数据，进行相加
			for job := range JobChan {
				// 随机数
				r_num := job.RandNum
				// 随机数每一位相加
				// 定义返回值
				var sum int
				for r_num != 0 {
					tmp := r_num % 10
					sum += tmp
					r_num /= 10
				}
				// 结果是Result
				r := &Result{job, sum}
				// 运算结果放入管道
				resultChan <- r
			}
		}(jobChan, resultChan)
	}
}
