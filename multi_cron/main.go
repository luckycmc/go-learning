package main

import (
	"fmt"
	"github.com/robfig/cron"
	"log"
)

type TestJob struct {
}

func (this TestJob) Run() {
	fmt.Println("hello world1")
}

type Test2Job struct {
}

func (this Test2Job) Run() {
	fmt.Println("hello world2")
}

func main() {
	i := 0
	c := cron.New()
	spec := "*/5 * * * * ?"
	c.AddFunc(spec, func() {
		i++
		log.Println("cron running: ", i)
	})
	c.AddJob(spec, TestJob{})
	c.AddJob(spec, Test2Job{})
	// 启动计划任务
	c.Start()
	// 关闭计划任务
	defer c.Stop()
	select {}
}
