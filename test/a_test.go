package test

import (
	"flag"
	"fmt"
	"os"
	"sync"
	"testing"

	"github.com/XZ0730/runFzu/pkg/utils"
)

func Test_A(t *testing.T) {
	var c = flag.Int("c", 1, "concurrency")
	var l = flag.Bool("l", false, "loop or not")
	var f *os.File
	var p = flag.String("p", "", "proxy url")
	flag.Parse()
	t.Log("----------------------")
	utils.ProxyURL = *p
	var wg sync.WaitGroup
	f, _ = os.Open("test.pcm")
	for i := 0; i < *c; i++ {
		fmt.Println("Main: Starting worker", i)
		wg.Add(1)
		if *l {
			go utils.ProcessLoop(i, &wg, f)
		} else {
			go utils.ProcessOnce(i, &wg, f)
		}
	}

	fmt.Println("Main: Waiting for workers to finish")
	wg.Wait()

	fmt.Println("Main: Completed")
}
