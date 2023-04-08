package materi

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func HttpGet() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
		if err != nil {
			log.Fatalln(err)
		}
	
		fmt.Println(resp.Body)
		body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
			log.Fatalln(err)
		}
	
		defer resp.Body.Close()
	
		sb := string(body)
		log.Println(sb)
}