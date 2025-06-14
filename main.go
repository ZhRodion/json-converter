package main

import (
	"fmt"
	"time"
)

type Bin struct {
	id        int
	private   bool
	createdAt time.Time
	name      string
}

type BinList struct {
	binList []Bin
}

func NewBin(id int, private bool, name string) *Bin {
	return &Bin{
		id:        id,
		private:   private,
		createdAt: time.Now(),
		name:      name,
	}
}

func (bl *BinList) AddBin(bin *Bin) {
	bl.binList = append(bl.binList, *bin)
}

func main() {
	binList := BinList{}

	bin := NewBin(1, true, "test")
	binList.AddBin(bin)

	fmt.Printf("%+v\n", binList)
}
