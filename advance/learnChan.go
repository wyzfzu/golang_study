package advance

import (
	"fmt"
	"time"
)

func produce(num int, ch chan<- int) {
	for i := 0; i < num; i++ {
		ch <- i
	}
}

func consume(ch chan int) {
	for {
		i, ok := <-ch
		if !ok {
			return
		}
		fmt.Println("receive ", i)
	}
}

func TestChan() {
	fmt.Println("chan with no buffer...")
	ch := make(chan int, 1)
	go consume(ch)
	go produce(10, ch)

	time.Sleep(2 * time.Second)

	close(ch)

	fmt.Println("chan with buffer...")
	ch = make(chan int, 100)

	go consume(ch)
	go produce(100, ch)

	time.Sleep(3 * time.Second)

	close(ch)
}
