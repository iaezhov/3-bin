package storage

import (
	"encoding/json"
	"hw/3/bins"
)

type Storage interface {
	Read() ([]byte, error)
	Write([]byte) error
}
type BinStorage struct {
	storage Storage
}

func NewBinStorage(storage Storage) *BinStorage {
	return &BinStorage{storage: storage}
}

func (bs *BinStorage) Save(list *bins.BinList) error {
	data, err := json.Marshal(list)
	if err != nil {
		return err
	}
	return bs.storage.Write(data)
}

func (bs *BinStorage) Load() (*bins.BinList, error) {
	data, err := bs.storage.Read()
	if err != nil {
		return nil, err
	}
	var list bins.BinList
	if err := json.Unmarshal(data, &list); err != nil {
		return nil, err
	}
	return &list, nil
}
