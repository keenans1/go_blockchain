package main

import (
	"bufio"
	"fmt"
	"os"
)

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
