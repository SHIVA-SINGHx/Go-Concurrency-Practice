package main

import (
	"fmt"
	"sync"
	"time"
)

// Without concurrency

// func worker(url string) string {

// 	time.Sleep(time.Microsecond * 50)

// 	fmt.Printf("image processed %s\n", url)
// 	return url
// }

// func main() {
// 	startTime:= time.Now()

// 	result1:=  worker("image1_png")
// 	result2 := worker("image2_png")

// 	worker("img_url1")

// 	fmt.Println("result1", result1)
// 	fmt.Println("result2", result2)

// 	fmt.Printf("it took %s ms.\n", time.Since(startTime))

// }

// GO Concurrency with goroutines

// create struct 
type Result struct{
	url string
	Err error
}

func worker(url string, wg *sync.WaitGroup , resultChan chan Result) {
	defer wg.Done() // it will run definitely when the function ends

	time.Sleep(time.Microsecond * 50)

	fmt.Printf("image processed %s\n", url)

	resultChan <- Result{
		url: url,
		Err: nil,
	}  // sending result to channel

}

func main(){

	var wg sync.WaitGroup
	resultChan:= make(chan Result, 5) // buffered channel with capacity 10

	startTime:= time.Now()

	wg.Add(5)
	go worker("img_1.png", &wg, resultChan)
	go worker("img_2.png", &wg, resultChan)
	go worker("img_3.png", &wg, resultChan)
	go worker("img_4.png", &wg, resultChan)
	go worker("img_5.png", &wg, resultChan)
	wg.Wait()
	close(resultChan) // close the channel after all goroutines are done

	for result:= range resultChan{
		fmt.Printf("recieved %v \n", result)  // reading from channel
	}

	fmt.Printf("it took %s ms.\n", time.Since(startTime))
}
