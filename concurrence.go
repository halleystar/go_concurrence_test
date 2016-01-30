package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

func main() {
	times := make(chan int, 3)
	go func() {
		for {
			times <- 1
			endTime := time.Now().Unix() + 3
			hits := 300
			mixHits := 1000
			//混合
			myChans := make(chan int, hits)
			b := ``
			url := ""

			f := func() {
				body := bytes.NewBuffer([]byte(b))
				response, err := http.Post(url, "application/json;charset=utf-8", body)
				if err != nil {
					fmt.Println(err)
				} else {
					if response.StatusCode != 200 {
						fmt.Println(response)
					}
				}
				myChans <- 1
			}
			for i := 0; i < hits; i++ {
				time.AfterFunc(time.Duration(endTime-time.Now().Unix()), f)
			}
			otherHits := mixHits - hits
			//TODO 完成混合请求
			go func() {
				time.Sleep(2 * time.Second)
			}()

			for i := 0; i < hits; i++ {
				<-myChans
			}
		}
		time.Sleep(3 * time.Second)
	}()
	for i := 0; i < 3; i++ {
		<-times
	}
}
