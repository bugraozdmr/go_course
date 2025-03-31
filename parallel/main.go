package main

import (
	"fmt"
	"time"
)

func greet(text string, doneChan chan bool) {
	fmt.Println("Hello ", text)
	doneChan <- true
}

func slowGreet(text string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello ", text)

	doneChan <- true
	// channel tamam olunca yenilerini bekleme diyoruz
	close(doneChan)
}

// basinda go olunca parallel oluyor - ancak bu bir şey dönmez
func main() {
	// channel - birden cok value alabilir
	// done := make(chan bool)
	// dones := make([]chan bool, 4)

	done := make(chan bool)

	// dones[0] = make(chan bool)
	go greet("Grant", done)
	// dones[1] = make(chan bool)
	go greet("The Emperor Grant Wick", done)

	// dones[2] = make(chan bool)
	go slowGreet("GGrant", done)
	// dones[3] = make(chan bool)
	go greet("Mark Greyson", done)
	//go greet("Mark Greyson", dones[3])

	/*for _, done := range dones {
		<-done
	}*/

	//<-done
	// fmt.Println(<- done)

	/*for doneChan := range done {
		fmt.Println(doneChan)
	}*/

	for range done {

	}
}
