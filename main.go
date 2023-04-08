package main

import (
	"challenge-9/ch9"
	"time"
)

func main() {
	ch9.PostJsonPlaceHolder()

	ticker := time.Tick(15 * time.Second)
	go func() {
		for range ticker {
			ch9.PostJsonPlaceHolder()
			
		}
	}()

	select{}
}

