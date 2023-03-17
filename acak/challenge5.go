package main

import (
	"fmt"
	"sync"
	"time"
)

type dataInterface interface{}
var wg sync.WaitGroup

func main() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()


	var interface1 dataInterface = []string{"coba1", "coba2", "coba3"}
	var interface2 dataInterface = []string{"bisa1", "bisa2", "bisa3"}
	for i := 1; i <= 4; i++ {
		wg.Add(2)
		go printXTimes(i ,4, interface1.([]string))
		go printXTimes(i ,4, interface2.([]string))
	}

	wg.Wait()	
}

func printXTimes(i int,num int, data []string) {
		fmt.Println(data, i)
		wg.Done()
}