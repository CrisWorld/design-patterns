package main

import (
	"fmt"
	"sync"
)

type HelloClass struct {
	Message string
}

type Singleton struct {
	instance *HelloClass
	once     sync.Once
}

// GetInstance cung cấp một instance duy nhất của HelloClass
func (s *Singleton) GetInstance() *HelloClass {
	// Sử dụng once.Do để khởi tạo s.instance chỉ một lần duy nhất
	s.once.Do(func() {
		s.instance = &HelloClass{Message: "Hello, Singleton!"}
	})
	return s.instance
}

func main() {
	singleton := &Singleton{}
	instance := singleton.GetInstance()
	instance2 := singleton.GetInstance()
	if instance == instance2 {
		fmt.Println("Same instance")
	}
	fmt.Println(instance.Message) // "Hello, Singleton!"
}
