package advance

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/**
题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ： sync.Mutex 的使用、并发数据安全。
题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ：原子操作、并发数据安全。
*/

var shareCount int = 0

func countWithMutex() {
	var m sync.Mutex
	var wg sync.WaitGroup

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				m.Lock()
				shareCount++
				m.Unlock()
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("shareCount结果: ", shareCount)
}

func countWithAtomic() {
	var cnt atomic.Int32
	var wg sync.WaitGroup

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				cnt.Add(1)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("atomicCount结果: ", cnt.Load())
}

func TestLock() {
	countWithMutex()
	countWithAtomic()
}
