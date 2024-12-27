package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Block struct {
	Index        int    
	Timestamp    string 
	Data         string 
	PreviousHash string 
	Hash         string 
}

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

func createAndSaveBlockchain(blockchain []Block) []Block{
	// Initialize Blockchain
	blockchain = append(blockchain, createGenesisBlock())
	// fmt.Println("Genesis Block Created:", blockchain[0])
	SaveBlockchain(blockchain)
	return blockchain
}

func ShowMenu() {
    fmt.Println("Choose an action:")
    fmt.Println("1. Add a new block")
    fmt.Println("2. View blockchain")
    fmt.Println("3. Verify blockchain")
    fmt.Println("4. Reset blockchain")
    fmt.Println("5. Exit")
}

func ResetBlockchain(blockchain []Block) []Block {
    if len(blockchain) == 0 {
        fmt.Println("Blockchain is empty; cannot reset.")
        return nil
    }
    // Retain only the first element (genesis block)
    return blockchain[:1]
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

func main() {
	blockChain, err := LoadBlockchain()

	if err != nil {
		blockChain = createAndSaveBlockchain(blockChain)
	} else {
		blockChain, _ = LoadBlockchain()
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		ShowMenu()
		var choice int
		fmt.Scanln(&choice)
	
		switch choice {
		case 1:
			fmt.Print("Enter data for the new block: ")
			scanner.Scan()
			input := scanner.Text()
			newBlock := generateNewBlock(blockChain[len(blockChain) - 1], input)
			blockChain = append(blockChain, newBlock)
			fmt.Println("Block added!")
		case 2:
			ViewBlockchain(blockChain)
		case 3:
			if isValidChain(blockChain) {
				fmt.Println("Blockchain is valid.")
			} else {
				fmt.Println("Blockchain integrity is compromised!")
			}
		case 4:
			blockChain = ResetBlockchain(blockChain)
			fmt.Println("Blockchain reset to genesis block.")
		case 5:
			fmt.Println("Exiting program.")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
