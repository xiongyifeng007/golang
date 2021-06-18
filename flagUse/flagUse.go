package main
 
import (
	"flag"
	"fmt"
)
 
//flag的作用是用来解析命令行的参数
// go run main.go -ip
func main(){
	//返回的是一个ip的指针
	ip := flag.String("ip","127.0.0.0","help message for flagname")
	flag.Parse()
	fmt.Printf("%v\n",*ip)
}
