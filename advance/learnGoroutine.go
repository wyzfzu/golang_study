package advance

import (
	"fmt"
	"sync"
	"time"
)

func printOdd() {
	for i := 1; i < 10; i += 2 {
		fmt.Print(i, ", ")
	}
	fmt.Println()
}

func printEven() {
	for i := 0; i < 10; i += 2 {
		fmt.Print(i, ", ")
	}
	fmt.Println()
}

func goPrint() {
	go printOdd()
	go printEven()
	time.Sleep(2 * time.Second)
}

type Task func()

var wg sync.WaitGroup

func scheduleTask(tasks []Task) {
	for _, task := range tasks {
		wg.Add(1)
		go func(task Task) {
			defer wg.Done()
			start := time.Now()
			task()
			end := time.Now()
			cost := end.Sub(start)

			fmt.Println("task ", task, "cost milli time: ", cost.Milliseconds())
		}(task)
	}
	wg.Wait()
}

func fastLoopTask() {
	time.Sleep(50 * time.Millisecond)
}

func slowLoopTask() {
	time.Sleep(1 * time.Second)
}

func normalLoopTask() {
	time.Sleep(200 * time.Millisecond)
}

func TestGoRutine() {
	goPrint()

	tasks := []Task{
		fastLoopTask, slowLoopTask, normalLoopTask,
	}

	scheduleTask(tasks)
}
