package bins

import (
	"encoding/json"
	"fmt"
)

type Vault struct {
	Bins []Bin `json:"bins"`
}

type Db interface {
	Read() ([]byte, error)
	Write([]byte) error
}

type VaultWithDb struct {
	Vault
	db Db
}

func NewVault(db Db) *VaultWithDb {
	file, err := db.Read()
	if err != nil {
		return &VaultWithDb{
			Vault: Vault{
				Bins: []Bin{},
			},
			db: db,
		}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	fmt.Printf("Найдено записей: %d\n", len(vault.Bins))
	if err != nil {
		fmt.Println(err)
		return &VaultWithDb{
			Vault: Vault{
				Bins: []Bin{},
			},
			db: db,
		}
	}
	return &VaultWithDb{
		Vault: vault,
		db:    db,
	}
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (vault *VaultWithDb) Save() {
	data, err := vault.ToBytes()
	if err != nil {
		fmt.Println("Не удалось преобразовать в JSON")
		return
	}
	vault.db.Write(data)
}

func (vault *VaultWithDb) AddBin(bin Bin) {
	vault.Bins = append(vault.Bins, bin)
	vault.Save()
}

func (vault *VaultWithDb) DeleteBin(id string) bool {
	var bins []Bin
	isDeleted := false
	for _, bin := range vault.Bins {
		isMatched := bin.ID == id
		if !isMatched {
			bins = append(bins, bin)
			continue
		}
		isDeleted = true
	}
	vault.Bins = bins
	vault.Save()
	return isDeleted
}
