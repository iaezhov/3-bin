package main

import (
	"fmt"
	"hw/3/bins"
	"hw/3/file"
	"hw/3/storage"
)

func main() {
	demoData := []string{"Первый", "Второй", "Третий"}

	binList := bins.NewBinList()
	for i, value := range demoData {
		binList.Add(bins.NewBin(fmt.Sprintf("%d", i), value, false))
	}

	filename := "bins.json"

	if !file.IsJSONFile(filename) {
		fmt.Println("Файл должен иметь расширение .json")
		return
	}

	if err := storage.SaveBins(filename, binList); err != nil {
		fmt.Println("Ошибка сохранения:", err)
		return
	}
	fmt.Println("Список сохранён в", filename)

	loadedList, err := storage.LoadBins(filename)
	if err != nil {
		fmt.Println("Ошибка загрузки:", err)
		return
	}

	outputResult(loadedList)
}

func outputResult(list *bins.BinList) {
	for _, b := range list.Bins {
		fmt.Println(b.GetName())
	}
}
