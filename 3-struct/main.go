package main

import (
	"fmt"
	"time"
)

type Bin struct {
	id       string
	private  bool
	createAt time.Time
	name     string
}

func createBin(id string, private bool, createAt time.Time, name string) Bin {
	return Bin{
		id:       id,
		private:  private,
		createAt: createAt,
		name:     name,
	}
}

func main() {
	bins := []Bin{}
	bin := createBin("asd", true, time.Now(), "qwer")
	fmt.Println(bin)

	for i := 0; i < 4; i++ {
		bins = append(bins, createBin("asd", true, time.Now(), "qwer"))
	}

	fmt.Println(bins)
}
