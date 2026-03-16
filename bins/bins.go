package bins

import "time"

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

func (bin *Bin) GetName() string {
	return bin.name
}

func NewBin(id string, name string, private bool) *Bin {
	return &Bin{
		id:        id,
		private:   private,
		name:      name,
		createdAt: time.Now(),
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
