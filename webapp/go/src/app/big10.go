package main

import "math/big"

var big10List = make([]*big.Int, 0, 10000)

func init() {
	for i := 0; i < 10000; i++ {
		tmp, _ := new(big.Int).SetString(big10[:i+1], 10)
		big10List = append(big10List, tmp)
	}

}
