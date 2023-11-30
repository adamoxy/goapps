package main

import (
	"fmt"
	"sync"
)

var (
	sharedBuffer []byte
	rwMutex      sync.RWMutex
)

func readFromBuffer(id int) {
	for {
		rwMutex.RLock()
		fmt.Printf("Reader %d: Reading data: %v\n", id, sharedBuffer)
		rwMutex.RUnlock()
	}
}

func writeToBuffer(id int, data byte) {
	for {
		rwMutex.Lock()
		fmt.Printf("Writer %d: Writing data: %d\n", id, data)
		sharedBuffer = append(sharedBuffer, data)
		rwMutex.Unlock()
	}
}

func main() {
	var m, n int

	fmt.Println("Enter M Value:")
	_, err := fmt.Scanln(&m)
	if err != nil {
		fmt.Println("Error reading M Value:", err)
		return
	}

	fmt.Println("Enter N value:")
	_, err = fmt.Scanln(&n)
	if err != nil {
		fmt.Println("Error reading N Value:", err)
		return
	}
	sharedBuffer = make([]byte, 0)

	// M readers
	for i := 0; i < m; i++ {
		go readFromBuffer(i)
	}
	// N writers
	for j := 0; j < n; j++ {
		go writeToBuffer(j, byte(j))
	}

	select {}
}
