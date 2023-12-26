package routine

import (
	"fmt"
	"sync"
)

func TestRoutine() {
	var wg sync.WaitGroup
	wg.Add(2)
	go multiply(&wg, 2, 4)
	go sum(&wg, 2, 3)
	wg.Wait()
}

func multiply(wg *sync.WaitGroup, a int, b int) {
	defer wg.Done()
	fmt.Println("testing", a*b)
}

func sum(wg *sync.WaitGroup, a int, b int) {
	defer wg.Done()
	fmt.Println("testing", a+b)
}
