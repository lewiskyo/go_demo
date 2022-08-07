package main

// https://www.codeleading.com/article/71415395181/
import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printAnimal(fromCh, toChan chan bool, printStr string, time int) {
	defer wg.Done()     // key point 1
	defer close(toChan) // key point 1

	for i := 0; i < time; i++ {
		<-fromCh
		fmt.Println(printStr)
		toChan <- true
	}
}

func main() {

	dogChan := make(chan bool, 1) // key point 3, buffer cap can not zero
	catChan := make(chan bool, 1)
	fishChan := make(chan bool, 1)
	dogChan <- true

	time := 4

	go printAnimal(dogChan, catChan, "dog...", time)
	go printAnimal(catChan, fishChan, "cat...", time)
	go printAnimal(fishChan, dogChan, "fish...", time)

	wg.Add(3)
	wg.Wait()
}
