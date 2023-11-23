package main

import (
	"fmt"
	"sync"
)

type SingleTon struct {
	Name string
}

var instance *SingleTon
var once sync.Once // sync.Once ensures that the initialization code is executed only once

func getInstance() *SingleTon {
	once.Do(func() {
		fmt.Println("CREATING SINGLETON INSTAMCE")
		instance = &SingleTon{Name: "data"}
	})
	return instance
}

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			singleton := getInstance()
			fmt.Printf("Goroutine %d: Singleton Data: %s\n", id, singleton.Name)
		}(i)

	}

	wg.Wait()

}
