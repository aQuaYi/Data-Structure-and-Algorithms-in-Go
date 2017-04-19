package main

import (
	"time"

	highcharts "github.com/aQuaYi/gohighcharts"
)

func main() {
	data := make(chan interface{})
	options := map[string]interface{}{
		"series": []interface{}{
			map[string]interface{}{
				"name": "Dynamic chart",
				"data": []int{},
			},
		},
		"chart": map[string]interface{}{
			"type": "line",
		},
	}
	highcharts.NewDynamicChart("/dynamic/", options, data)
	go func() {
		for i := 0; i < 10; i++ {
			data <- i
			time.Sleep(time.Second * 3)
		}
	}()
	time.Sleep(time.Second * 100000)
}
