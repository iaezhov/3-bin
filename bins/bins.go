package bins

import "time"

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	Name      string
}

func NewBin(id string, name string, private bool) *Bin {
	return &Bin{
		id:        id,
		private:   private,
		Name:      name,
		createdAt: time.Now(),
	}
}

type BinList []Bin
