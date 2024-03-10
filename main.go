package main

import (
	"AnimeImageDownloader/service"
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {
	_ = os.Mkdir("image", 0777)
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
			}()
			service.Down("./image/")
		}()
	}
	wg.Wait()
}
