package storage

import (
	"encoding/json"
	"hw/3/bins"
	"hw/3/file"
)

func SaveBins(filename string, list *bins.BinList) error {
	data, err := json.Marshal(list)
	if err != nil {
		return err
	}
	_, err = file.WriteFile(data, filename)
	return err
}

func LoadBins(filename string) (*bins.BinList, error) {
	data, err := file.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var list bins.BinList
	if err := json.Unmarshal(data, &list); err != nil {
		return nil, err
	}
	return &list, nil
}
