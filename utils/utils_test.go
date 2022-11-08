package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestNone(t *testing.T) {

	ch := make(chan int, 2)

	go func(ch chan int) {
		for v := range ch {
			fmt.Println("process gain val", v)
			time.Sleep(time.Second * 2)
		}
	}(ch)

	for i := 0; i < 10; i++ {
		select {
		case ch <- i:
			fmt.Println("push to chan", i)
		default:
			fmt.Println("ignore val", i)
		}
		time.Sleep(time.Second)
	}
}
