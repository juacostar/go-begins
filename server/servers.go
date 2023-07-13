package main

import (
	"fmt"
	"net/http"
	"time"
)


func main(){
	begin := time.Now()

	channel := make(chan string) // channel with data types that pass the channel

	servers := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebbok.com",
		"http://instagram.com",
	}

	for _, server := range servers {
		go seeServer(server, channel) // launching goroutines with go
		//fmt.Println(<- channel) sequential way
	}

	for i:=0 ; i < len(servers); i++ {
		fmt.Println(<- channel) // receiving channel information
	}

	duration := time.Since(begin)
	fmt.Println(duration)
}

func seeServer(server string, channel chan string){
	_, err := http.Get(server)
	if err != nil {
		fmt.Println("error")
		channel <- "error"
	}else{
		fmt.Println("done")
		channel <- "done" // set value to eh channel
	}
}