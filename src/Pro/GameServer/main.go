/*
* CopyRight(C) StefanChen e-mail:2572915286@qq.com
*/

package main

import (
	"fmt"
	"flag"
	//"log"
)


func init(){

}

var addr = flag.String("addr", "localhost:8080", "http service address")

func main()  {
	fmt.Println("start game.")
	flag.Parse()
	
}