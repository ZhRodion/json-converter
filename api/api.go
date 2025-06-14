package api

import (
	"fmt"
	"json-converter/bins"
)

func Api(binList *bins.BinList) {

	bin := bins.NewBin(1, true, "test")
	binList.AddBin(bin)

	fmt.Printf("%+v\n", binList)
}
