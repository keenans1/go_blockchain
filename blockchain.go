package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func isValidChain(blockchain []Block) bool {
	for i := 1; i < len(blockchain); i++ {
		currentBlock := blockchain[i]
		previousBlock := blockchain[i-1]

		if currentBlock.PreviousHash != previousBlock.Hash {
			return false
		}
		if currentBlock.Hash != calculateHash(currentBlock) {
			return false
		}
	}
	return true
}

func SaveBlockchain(chain []Block) error {
    data, err := json.Marshal(chain)
    if err != nil {
        return err
    }
    return os.WriteFile("blockchain.json", data, 0644)
}

func LoadBlockchain() ([]Block, error) {
    data, err := os.ReadFile("blockchain.json")
    if err != nil {
        return nil, err
    }
    var chain []Block
    json.Unmarshal(data, &chain)
    return chain, nil
}

func createAndSaveBlockchain(blockchain []Block) []Block{
	// Initialize Blockchain
	blockchain = append(blockchain, createGenesisBlock())
	// fmt.Println("Genesis Block Created:", blockchain[0])
	SaveBlockchain(blockchain)
	return blockchain
}

func ResetBlockchain(blockchain []Block) []Block {
    if len(blockchain) == 0 {
        fmt.Println("Blockchain is empty; cannot reset.")
        return nil
    }
    // Retain only the first element (genesis block)
    return blockchain[:1]
}
