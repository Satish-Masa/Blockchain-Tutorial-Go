package main

import "fmt"

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 BTC to jack")
	bc.AddBlock("Send 2 more BTC to Jack")

	for _, block := range bc.blocks {
		fmt.Printf("Prev hash:             %x\n", block.PreBlockHash)
		fmt.Printf("Data                   %s\n", block.Data)
		fmt.Printf("Hash                   %x\n", block.Hash)
		fmt.Println()
	}
}
