package main

import (
	"essential/essential"
	"fmt"
)

func main() {
	chain := essential.NewBlockchain("Genesis Information")
	chain.AddBlock("Data 1")
	chain.AddBlock("Data 2")
	chain.AddBlock("Data 3")
	for _, block := range chain.Blocks {
		fmt.Println(block)
		fmt.Println("=====")
	}
}
