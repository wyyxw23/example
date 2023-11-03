package main

import (
	"example.com/greetings"
	"fmt"
	"log"
)

func main() {
	// 设置预定义的logger属性，
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	// 获取问候信息并打印出来.
	names := []string{"Gladys", "Samantha", "Darrin"}
	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
