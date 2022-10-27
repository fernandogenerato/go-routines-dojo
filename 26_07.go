package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

func slowfunction(from string) error {
	if from == "" {
		panic("panic")
		return errors.New("slow error")
	}
	fmt.Println(from)
	time.Sleep(2 * time.Second)
	return nil
}

func main() {
	wg := sync.WaitGroup{}

	wg.Add(3)
	direct := []string{
		"direct1",
		"",
		"direct3",
	}

	for _, s := range direct {
		go func(str string) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println(r)
					wg.Done()
				}
			}()

			err := slowfunction(str)


			if err != nil {
				return
			}
			wg.Done()

		}(s)
	}
	wg.Wait()
	fmt.Println("done")
}
