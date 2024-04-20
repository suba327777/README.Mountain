package main

import "fmt"

func getHello() string {
	return "hello world"
}

func typo() {
	mispell := "misspell"
	println(mispell)
}

func main() {
	fmt.Println(getHello())
}
