package main

import (
	"fmt"
	"hw/3/api"
	"hw/3/bins"
	"hw/3/config"
	"hw/3/file"
	"hw/3/storage"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Ошибка загрузки ENV")
		return
	}

	api.OutputConfig(config.NewConfig())

	// Наполнение данными
	demoData := []string{"Первый", "Второй", "Третий"}
	binList := bins.NewBinList()
	for i, value := range demoData {
		binList.Add(bins.NewBin(fmt.Sprintf("%d", i), value, false))
	}

	// запись данных
	filename := "bins.json"
	fileStorage, err := file.NewJSONFileStorage(filename)
	if err != nil {
		fmt.Println("Ошибка создания хранилища:", err)
		return
	}

	binStorage := storage.NewBinStorage(fileStorage)
	if err := binStorage.Save(binList); err != nil {
		fmt.Println("Ошибка сохранения:", err)
		return
	}
	fmt.Println("Список сохранён в", filename)

	// чтение данных
	loadedList, err := binStorage.Load()
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
