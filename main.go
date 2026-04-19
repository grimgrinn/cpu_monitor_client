package main

import (
	"io"
	"fmt"
	"net/http"
	"time"
)

func main() {
	for {
		resp, err := http.Get("http://188.214.144.76:8080/load")
		if err != nil {
			fmt.Printf("error: %v \n", err)
			return
		}
			
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close() 
		if err != nil {
			fmt.Printf("error reading: %v\n", err)
			return
		}
	
		fmt.Printf("\r%s", string(body))

		time.Sleep(1000 * time.Millisecond)
		//fmt.Print("\033[1A\033[K")
	}
}

