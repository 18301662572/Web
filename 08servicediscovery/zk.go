package main

import (
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

//基于 zk 的完整服务发现流程

func main() {
	c, _, err := zk.Connect([]string{"127.0.0.1"}, time.Second)
	if err != nil {
		panic(err)
	}
	res, err := c.Create("/platform/order-system/create-order-service-http/10.1.13.3:1043", []byte("1"),
		zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	if err != nil {
		panic(err)
	}
	println(res)
	time.Sleep(time.Second * 50)
}

//在 sleep 的时候我们在 cli 中查看写入的临时节点数据：
//
//ls /platform/order-system/create-order-service-http
//['10.1.13.3:1043']
//在程序结束之后，很快这条数据也消失了：
//
//ls /platform/order-system/create-order-service-http
//[]
