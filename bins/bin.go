package bins

import "time"

type Bin struct {
	ID        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

func NewBin(id string, name string, private bool) *Bin {
	return &Bin{
		ID:        id,
		Private:   private,
		Name:      name,
		CreatedAt: time.Now(),
	}
}
