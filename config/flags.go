package config

import (
	"flag"
)

type Flags struct {
	FileName     string
	BinId        string
	BinName      string
	ActionCreate bool
	ActionUpdate bool
	ActionDelete bool
	ActionGet    bool
	Actionlist   bool
}

func NewFlags() *Flags {
	fileName := flag.String("file", "bins.json", "File name")
	binId := flag.String("id", "", "Bin ID")
	binName := flag.String("name", "", "Bin name")
	actionCreate := flag.Bool("create", false, "Create Bin")
	actionUpdate := flag.Bool("update", false, "Update Bin")
	actionDelete := flag.Bool("delete", false, "Delete Bin")
	actionGet := flag.Bool("get", false, "Get Bin")
	actionList := flag.Bool("list", false, "Get Bins")
	flag.Parse()

	return &Flags{
		BinId:        *binId,
		FileName:     *fileName,
		BinName:      *binName,
		ActionCreate: *actionCreate,
		ActionUpdate: *actionUpdate,
		ActionDelete: *actionDelete,
		ActionGet:    *actionGet,
		Actionlist:   *actionList,
	}
}
