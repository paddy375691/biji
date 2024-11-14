package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	// add方法
	fmt.Println("当前时间", now)
	m, _ := time.ParseDuration("-10m")
	m1 := now.Add(m)
	fmt.Println("前一分钟", m1)
	// sub
}
