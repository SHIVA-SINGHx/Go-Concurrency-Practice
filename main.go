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

func worker(url string, wg *sync.WaitGroup ){
	defer wg.Done() // it will run definitely when the function ends

	time.Sleep(time.Microsecond * 50)

	fmt.Printf("image processed %s\n", url)


}

func main(){

	var wg sync.WaitGroup

	startTime:= time.Now()

	wg.Add(10)
	go worker("img_1.png", &wg)
	go worker("img_2.png", &wg)
	go worker("img_3.png", &wg)
	go worker("img_4.png", &wg)
	go worker("img_5.png", &wg)
	go worker("img_6.png", &wg)
	go worker("img_7.png", &wg)
	go worker("img_8.png", &wg)
	go worker("img_9.png", &wg)
	go worker("img_10.png", &wg)
	wg.Wait()

	fmt.Printf("it took %s ms.\n", time.Since(startTime))
}
