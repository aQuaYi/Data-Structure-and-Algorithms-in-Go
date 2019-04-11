package counter

import (
	"fmt"
	"os"
	"sync"
)

//Counter 是计数器
type VisualCounter struct {
	n, max, number int
	mutex          sync.Mutex
}

//New 返回一个*Counter实例
func New(n, max int) *VisualCounter {
	return &VisualCounter{
		n:   n,
		max: max,
	}
}

//Increment 添加一个数字
func (c *VisualCounter) Increment() {
	c.mutex.Lock()
	c.number++
	c.n--
	c.mutex.Unlock()
	c.checkN()
	c.checkMax()
}

//Decrement 减小计数器一个数字
func (c *VisualCounter) Decrement() {
	c.mutex.Lock()
	c.number--
	c.n--
	c.mutex.Unlock()
	c.checkN()
	c.checkMax()
}

func (c *VisualCounter) checkN() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.n < 0 {
		fmt.Println("统计次数用完")
		os.Exit(1)
	}
}

func (c *VisualCounter) checkMax() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.number > c.max {
		fmt.Println("超过计数器极限")
		os.Exit(1)
	}
}

//Tally 返回计数器的次数
func (c *VisualCounter) Tally() int {
	return c.number
}

func (c *VisualCounter) String() string {
	return fmt.Sprintf("总计数是%d", c.number)
}
