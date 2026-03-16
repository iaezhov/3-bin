package main

import (
	"fmt"
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

func newBin(id string, name string, private bool) *Bin {
	return &Bin{
		id:        id,
		private:   private,
		name:      name,
		createdAt: time.Now(),
	}
}

type BinList []Bin

func main() {
	demoData := []string{"Первый", "Второй", "Третий"}

	var binList = make(BinList, len(demoData))

	for i, value := range demoData {
		binList[i] = *newBin(fmt.Sprintf("%d", i), value, false)
	}

	outputResult(&binList)
}

func outputResult(list *BinList) {
	for _, value := range *list {
		fmt.Println(value.name)
	}
}
