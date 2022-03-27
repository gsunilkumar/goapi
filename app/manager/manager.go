package manager

import (
	"fmt"
	"strconv"
	"time"
)

func MakeRequest(url string, stop chan bool) {
	fmt.Println("Processing request", url)
	time.Sleep(3 * time.Second)
	fmt.Println("Finished pre-processing request", url)

	for {
		select {
		case <-stop:
			fmt.Println("Stopping")
			return
		default:
		}
	}

	fmt.Println("Processed request", url)
}


func Process(){
	stop := make(chan bool)
	for i := 1; i < 5; i++ {
		go MakeRequest(strconv.Itoa(i), stop)
	}

	time.Sleep(1*time.Second)

	close(stop) // cancel current goroutines
}