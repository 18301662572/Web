package main
import (
	"fmt"
	"time"
)

//令牌桶demo--令牌桶模型原理

//github.com/juju/ratelimit 库：
//代码里向 channel 里填充 token 的操作，理论上是没有必要的。只要在每次 Take 的时候，再对令牌桶中的 token 数进行简单计算，
// 就可以得到正确的令牌数.
//在得到正确的令牌数之后，再进行实际的 Take 操作就好，这个 Take 操作只需要对令牌数进行简单的减法即可，记得加锁以保证并发安全。
//cur = k1 + ((t2 - t1)/ti) * x
//cur = cur > cap ? cap : cur


var tokenBucket chan struct{}

//放令牌
func main() {
	var fillInterval = time.Millisecond * 10
	var capacity = 100
	tokenBucket = make(chan struct{}, capacity)
	fillToken := func() {
		//NewTicker 返回一个新的 Ticker，该 Ticker 包含一个通道字段，并会每隔时间段 d 就向该通道发送当时的时间。它会调
		//整时间间隔或者丢弃 tick 信息以适应反应慢的接收者。如果d <= 0会触发panic。关闭该 Ticker 可
		//以释放相关资源
		ticker := time.NewTicker(fillInterval)
		for {
			select {
			case <-ticker.C:
				select {
				case tokenBucket <- struct{}{}:
				default:
				}
				fmt.Println("current token cnt:", len(tokenBucket), time.Now())
			}
		}
	}
	go fillToken() //放令牌
	go TakeAvailable(true)  //取令牌
	time.Sleep(time.Hour)
}
//time.NewTicker定时触发执行任务，当下一次执行到来而当前任务还没有执行结束时，会等待当前任务执行完毕后再执行下一次任务。
//time.NewTimer和Reset()函数实现定时触发，Reset()函数可能失败，经测试。

//取令牌
func TakeAvailable(block bool) bool{
	var takenResult bool
	if block {
		select {
		case <-tokenBucket:
			takenResult = true
		}
	} else {
		select {
		case <-tokenBucket:
			takenResult = true
		default:
			takenResult = false
		}
	}
	return takenResult
}