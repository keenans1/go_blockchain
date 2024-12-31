package blockchain

import (
	"encoding/json"
	"fmt"
	"os"
)

func IsValidChain(blockchain []Block) bool {
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

func SaveBlockchain(blockchain []Block) error {
    data, err := json.Marshal(blockchain)
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
    var blockchain []Block
    json.Unmarshal(data, &blockchain)
    return blockchain, nil
}

func CreateAndSaveBlockchain(blockchain *[]Block) {
    *blockchain = append(*blockchain, createGenesisBlock())
    SaveBlockchain(*blockchain)
}

func ResetBlockchain(blockchain *[]Block) {
    if len(*blockchain) == 0 {
        fmt.Println("Blockchain is empty; cannot reset.")
        return
    }
    *blockchain = (*blockchain)[:1]
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
