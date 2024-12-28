package main

import (
	"fmt"
)

func ShowMenu() {
    fmt.Println("Choose an action:")
    fmt.Println("1. Add a new block")
    fmt.Println("2. View blockchain")
    fmt.Println("3. Verify blockchain")
    fmt.Println("4. Reset blockchain")
    fmt.Println("5. Exit")
}

func ViewBlockchain(blockChain []Block) {
	fmt.Println("Blockchain contents:")
    for _, block := range blockChain {
        fmt.Printf("Block %d:\n", block.Index)
        fmt.Printf("  Timestamp: %s\n", block.Timestamp)
        fmt.Printf("  Data: %s\n", block.Data)
        fmt.Printf("  PrevHash: %s\n", block.PreviousHash)
        fmt.Printf("  Hash: %s\n", block.Hash)
        fmt.Println("-------------------------")
    }
}
