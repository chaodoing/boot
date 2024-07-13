package tests

import (
	`fmt`
	`sync`
	`testing`
	`time`
)

func TestSync(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			func(v int) {
				fmt.Printf("Worker %d starting\n", v)
				time.Sleep(time.Second * 5)
				fmt.Printf("Worker %d done\n", v)
			}(i)
		}(i)
	}
	wg.Wait()
}
