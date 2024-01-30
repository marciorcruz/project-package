package main

import (
	"custompackage/contracts/erc1155"
	"custompackage/contracts/erc20"
	"fmt"
)

func main() {
	fmt.Println("Meu projeto")
	// custompackage.AbiERC()
	// custompackage.Lins("Luiz")
	// fmt.Println(custompackage.Soma(2, 2))
	erc20.Interface()
	erc1155.Interface()
}
