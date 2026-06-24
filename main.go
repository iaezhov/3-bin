package main

import (
	"encoding/json"
	"fmt"
	"hw/3/bins"
	"hw/3/config"
	"hw/3/file"
)

func main() {
	envs, err := config.NewEnvs()
	if err != nil {
		fmt.Println("Ошибка загрузки ENV")
		return
	}

	db, err := file.NewJSONFileStorage(envs.StoragePath)
	if err != nil {
		fmt.Println("Ошибка создания хранилища:", err)
		return
	}

	vault := bins.NewVault(db)
	api := bins.NewJsonBinApi(envs.ApiUrl, envs.ApiKey)
	flags := config.NewFlags()

	if flags.Actionlist {
		fmt.Println("ID", "|", "Name")
		fmt.Println("----------------")
		for _, b := range vault.Bins {
			fmt.Println(b.ID, "|", b.Name)
		}
		return
	}

	if flags.ActionGet {
		body, err := api.Get(flags.BinId)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("GET: Success")
		fmt.Println("--------")
		fmt.Println(string(body))
	}

	if flags.ActionUpdate {
		fileStorage, err := file.NewJSONFileStorage(flags.FileName)
		if err != nil {
			fmt.Println(err)
			return
		}

		data, err := fileStorage.Read()
		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = api.Update(flags.BinId, data)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("UPDATE: Success")
	}

	if flags.ActionDelete {
		_, err = api.Delete(flags.BinId)
		if err != nil {
			fmt.Println(err)
			return
		}
		success := vault.DeleteBin(flags.BinId)
		if !success {
			fmt.Println("Не удалось удалить запись из хранилища")
			return
		}
		fmt.Println("DELETE: Success")
	}

	if flags.ActionCreate {
		fileStorage, err := file.NewJSONFileStorage(flags.FileName)
		if err != nil {
			fmt.Println(err)
			return
		}
		data, err := fileStorage.Read()
		if err != nil {
			fmt.Println(err)
			return
		}

		body, err := api.Create(data)
		var response bins.BinResponse
		err = json.Unmarshal(body, &response)
		if err != nil {
			fmt.Println(err)
			return
		}
		vault.AddBin(response.Metadata)
		fmt.Println("CREATE: Success")
	}

	fmt.Println("Action required. Commands: create|get|delete|update|list")
}
