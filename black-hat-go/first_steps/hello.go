package main

import (
	"fmt"
	"net"
	"sync"
)

func greet() {
	var x = "hello old man"
	fmt.Println(x)
	fmt.Println("Hello, white Hat Gophers!")
}

func scan() {
	var wg sync.WaitGroup
	_, err := net.Dial("tcp", "scanme.nmap.org:80")
	if err == nil {
		fmt.Println("connection successful")
	}

	// Commented out as 1024 without a work pool gets the network pretty stuck
	// for i := 1; i <= 1024; i++ {
	for i := 1; i <= 23; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				// port is closed of filtered
				// fmt.Printf("%d no response\n", j)
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
	}
	wg.Wait()
}

func main() {
	greet()
	scan()
}
