package main

import (
	"fmt"
	"github.com/sourcegraph/conc"
	"time"
)

func main() {
	var wg conc.WaitGroup
	defer wg.Wait()

	for i := 0; i < 10; i++ {
		wg.Go(func() { fmt.Println("hello world" + time.Now().String()) })

	}
}
