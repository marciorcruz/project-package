package main

import (
	"custompackage"
	"fmt"
)

func main() {
	fmt.Println("Meu projeto")
	custompackage.AbiERC()
	custompackage.Lins("Luiz")
	fmt.Println(custompackage.Soma(2, 2))
}
