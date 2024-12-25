package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Block struct {
	Index        int    // Position in the chain
	Timestamp    string // Creation time
	Data         string // Block data
	PreviousHash string // Hash of the previous block
	Hash         string // Hash of the current block
}

// var Blockchain []Block

func calculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + block.Data + block.PreviousHash
	hash := sha256.Sum256([]byte(record))
	return fmt.Sprintf("%x", hash)
}

func createGenesisBlock() Block {
	genesisBlock := Block{
		Index:        0,
		Timestamp:    time.Now().String(),
		Data:         "Genesis Block",
		PreviousHash: "",
	}
	genesisBlock.Hash = calculateHash(genesisBlock)
	return genesisBlock
}

func generateNewBlock(previousBlock Block, data string) Block {
	newBlock := Block{
		Index:        previousBlock.Index + 1,
		Timestamp:    time.Now().String(),
		Data:         data,
		PreviousHash: previousBlock.Hash,
	}
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}

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

// func createAndLoadBlockchain(blockchain []Block) {
// 	// Initialize Blockchain
// 	Blockchain = append(Blockchain, createGenesisBlock())
// 	fmt.Println("Genesis Block Created:", Blockchain[0])

// 	// Add New Blocks
// 	Blockchain = append(Blockchain, generateNewBlock(Blockchain[len(Blockchain)-1], "First Block"))
// 	Blockchain = append(Blockchain, generateNewBlock(Blockchain[len(Blockchain)-1], "Second Block"))
// }

func createAndLoadBlockchain(blockchain []Block) []Block{
	// Initialize Blockchain
	blockchain = append(blockchain, createGenesisBlock())
	fmt.Println("Genesis Block Created:", blockchain[0])

	// Add New Blocks
	blockchain = append(blockchain, generateNewBlock(blockchain[len(blockchain)-1], "First Block"))
	blockchain = append(blockchain, generateNewBlock(blockchain[len(blockchain)-1], "Second Block"))

	// load into file
	SaveBlockchain(blockchain)

	return blockchain
}

func main() {

	blockChain, err := LoadBlockchain()

	if err != nil {
		// create and load
		blockChain = createAndLoadBlockchain(blockChain)
	} else {
		// load it
		blockChain, _ = LoadBlockchain()
	}

	// Print Blockchain
	// for _, block := range Blockchain {
	// 	fmt.Printf("Index: %d, Data: %s, Hash: %s\n", block.Index, block.Data, block.Hash)
	// }

	// Validate Chain
	if isValidChain(blockChain) {
		fmt.Println("Blockchain is valid!")
	} else {
		fmt.Println("Blockchain is invalid!")
	}
}
