package chmerge

import "sync"

func Merge[T any](channels ...chan T) <-chan T {
	out := make(chan T)
	var wg sync.WaitGroup
	wg.Add(len(channels))

	for _, c := range channels {
		ci := c
		go func() {
			for v := range ci {
				out <- v
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
