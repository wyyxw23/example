package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello 为指定的人返回问候语.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	// 使用随机格式创建消息
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	// 将名称与消息关联的map
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat 返回一组问候消息中的一个，返回的消息是随机选择的
func randomFormat() string {
	// 消息格式的切片
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	// 通过为格式切片指定随机索引
	// 来返回随机选择的消息格式
	return formats[rand.Intn(len(formats))]
}
