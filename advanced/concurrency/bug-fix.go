package concurrency

import (
	"fmt"
)

func BugFix() {
	obj := Counter{}
	cntrCh := make(chan int)
	for i := 1; i <= 1000; i++ {
		go obj.Count(cntrCh)
		obj.c = <-cntrCh
	}

	fmt.Println(obj.c)
}

type Counter struct {
	c int
}

func (ct *Counter) Count(coCh chan int) {
	ct.c++
	coCh <- ct.c
}
