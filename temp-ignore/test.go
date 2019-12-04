package main

import (
	"net/http"
	"strings"
	"fmt"
	"time"
)

func main(){
	body := strings.NewReader("Hi how are you")
	for i :=0; i< 100; i++ {
		go func(){
			res, err := http.Post("http://localhost:8080", "application/json", body)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(res)
			}
		}()
	}

	time.Sleep(10* time.Second)
}
