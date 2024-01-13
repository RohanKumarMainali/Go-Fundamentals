package routine

import (
	"fmt"
	"sync"
	"time"
)

func Concurrency() {
	now := time.Now()
	wg := &sync.WaitGroup{}
	buffer := 3
	responseChannel := make(chan string, buffer)
	go fetchUserData(responseChannel, wg)
	go fetchProducts(responseChannel, wg)
	go fetchGraph(responseChannel, wg)

	wg.Add(3)
	wg.Wait()
	close(responseChannel)
	for res := range responseChannel {
		fmt.Println(res)
	}

	fmt.Println(time.Since(now))
}

func fetchUserData(response chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(60 * time.Millisecond)
	response <- "rohan"
}

func fetchProducts(response chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(100 * time.Millisecond)
	response <- "products"
}

func fetchGraph(response chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(100 * time.Millisecond)
	response <- "graphs"
}
