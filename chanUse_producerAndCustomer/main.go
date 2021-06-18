//channel实现生产消费者模式
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(name string, channel chan<- string) { //生产者只能向通道输入数据 i := 0; i < 10; i++
	for i := 0; i < 10; i++ {
		channel <- fmt.Sprintf("%s:%v", name, rand.Int31())
		time.Sleep(time.Second)
	}
	close(channel) //通道关闭后,无法写入数据
	fmt.Println("channel closed")
}

func customer(channel <-chan string) { //消费者只能从通道接受数据
	for data := range channel { //通道关闭后,循环会结束.若不手动close通道将会报错
		fmt.Println(data)
	}
	fmt.Println("break for")
}

func main() {
	channel := make(chan string)
	go producer("Jack", channel)
	go producer("Tom", channel)
	customer(channel)
}
