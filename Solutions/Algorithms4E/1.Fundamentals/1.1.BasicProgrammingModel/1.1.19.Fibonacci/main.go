//
package main

import (
	"fmt"
	"sync"
	"time"
)

var f map[int]int
var mutex sync.Mutex

func main() {
	f = make(map[int]int)
	beginTime := time.Now()
	for i := 0; i < 50; i++ {
		beginTime = time.Now()
		fmt.Printf("Fast(%d)=%d,耗时%s\n", i, Fast(i), time.Since(beginTime))
		beginTime = time.Now()
		fmt.Printf("   F(%d)=%d,耗时%s\n", i, F(i), time.Since(beginTime))
	}
}

func F(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return F(n-1) + F(n-2)
}

func Fast(n int) (result int) {
	if v, ok := f[n]; ok {
		result = v
		return
	}
	switch n {
	case 0:
		result = 0
	case 1:
		result = 1
	default:
		result = Fast(n-1) + Fast(n-2)
	}
	{
		mutex.Lock()
		f[n] = result
		mutex.Unlock()
	}
	return
}
