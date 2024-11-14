package main

import (
	"fmt"
	"sync"
)

func main() {
	wg.Add(1)
	go getip(jobchan)
	wg.Add(4)
	//开启4个goroutine从jobchan中的ip取出
	for i := 0; i < 4; i++ {
		go PrintTotal(jobchan)
	}
	//time.Sleep(1000)
	wg.Wait()
}

type job struct {
	ip string
}

var jobchan = make(chan *job, 100)

//var resultchan = make(chan *result, 1)
var wg sync.WaitGroup

//1.开启一个goroutine循取出ip，发送到jobChan
func getip(ipchan chan<- *job) {
	defer wg.Done()
	ip_list := [...]string{
		"1111",
		"11231",
		"asdfcx",
		"sdaf",
		"xzcv",
		"jkljk",
		"xczxcv",
	}
	for _, v := range ip_list {
		y := v
		newjob := &job{
			ip: y,
		}
		ipchan <- newjob
	}
	close(ipchan)
}

//开启4个goroutine从jobchan中的ip取出并获取整体的api数据
func PrintTotal(ipchan <-chan *job) {
	defer wg.Done()
	for {
		j, ok := <-ipchan
		if !ok {
			break
		}
		cc(j.ip)
	}

}
func cc(ip string) {
	fmt.Println(ip)
}
