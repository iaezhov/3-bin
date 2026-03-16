package main

import (
	"fmt"
	"hw/3/bins"
	"hw/3/file"
)

func main() {
	demoData := []string{"Первый", "Второй", "Третий"}

	var binList = make(bins.BinList, len(demoData))

	for i, value := range demoData {
		binList[i] = *bins.NewBin(fmt.Sprintf("%d", i), value, false)
	}

	file.WriteFile()
	file.ReadFile()

	outputResult(&binList)
}

func outputResult(list *bins.BinList) {
	for _, value := range *list {
		fmt.Println(value.Name)
	}
}
