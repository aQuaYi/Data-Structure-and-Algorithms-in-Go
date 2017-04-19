package main

import "math/big"
import "fmt"

func main() {
	ba := big.NewFloat(2)
	bb := big.NewFloat(3)
	bc := big.NewFloat(5)
	bd := ba.Add(bb, bc)
	fmt.Println(bd)

}
