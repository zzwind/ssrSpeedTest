package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func SpeedTest(s ServerList, ms int) ServerList {
	aServerList := ServerList{}
	for _, v := range s {
		sp := strconv.Itoa(v.ServerPort)
		if connect(ms, v.Server, sp) {
			aServerList = append(aServerList, v)

		}
	}

	if len(aServerList) > 5 {
		fmt.Printf("可用的IP数量:%d,链接时间%d毫秒以内", len(aServerList), ms)
	} else {
		fmt.Println("开始新的测试")
		SpeedTest(s, ms+100)
	}
	return aServerList

}

func connect(t int, address string, port string) bool {
	_, err := net.DialTimeout("tcp", address+":"+port, time.Millisecond*time.Duration(t))
	if err != nil {
		return false
	}
	return true
}
