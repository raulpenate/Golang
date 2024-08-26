package main

import (
	"fmt"
	"sync"
)

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}

func main() {
	var wg sync.WaitGroup

	fmt.Println("Hello")
	wg.Add(2)

	go say("world", &wg)

	go func(text string, wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println(text)
	}("Adios", &wg)

	wg.Wait()

	text := "你好，世界" // "Hello, World" en chino
	fmt.Println(text)

}
