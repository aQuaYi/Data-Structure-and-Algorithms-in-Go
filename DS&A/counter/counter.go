package counter

import (
	"fmt"
	"sync"
)

//Counter 是计数器
type Counter struct {
	name   string
	number int
	mutex  sync.Mutex
}

//New 返回一个*Counter实例
func New(name string) *Counter {
	return &Counter{
		name: name,
	}
}

//Increment 添加一个数字
func (c *Counter) Increment() {
	c.mutex.Lock()
	c.number++
	c.mutex.Unlock()
}

//Tally 返回计数器的次数
func (c *Counter) Tally() int {
	return c.number
}

func (c *Counter) String() string {
	return fmt.Sprintf("总计数是%d", c.number)
}
