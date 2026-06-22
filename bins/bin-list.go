package bins

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
