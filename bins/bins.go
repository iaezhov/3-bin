package bins

import "time"

type Bin struct {
	ID        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

func (bin *Bin) GetName() string {
	return bin.Name
}

func NewBin(id string, name string, private bool) *Bin {
	return &Bin{
		ID:        id,
		Private:   private,
		Name:      name,
		CreatedAt: time.Now(),
	}
}

type Bins = []*Bin
type BinList struct {
	Bins Bins
}

func NewBinList() *BinList {
	return &BinList{
		Bins: make(Bins, 0),
	}
}

func (bl *BinList) Add(bin *Bin) {
	bl.Bins = append(bl.Bins, bin)
}
