package main

import (
	"bufio"
	"fmt"
	"os"
	"go_blockchain/user_input"
	"go_blockchain/blockchain"
)

func main() {
	blockChain, err := blockchain.LoadBlockchain()

	if err != nil {
		blockchain.CreateAndSaveBlockchain(&blockChain)
	} 
	
	scanner := bufio.NewScanner(os.Stdin)

	for {
		user_input.ShowMenu()
		var choice int
		fmt.Scanln(&choice)
	
		switch choice {
		case 1:
			fmt.Print("Enter data for the new block: ")
			scanner.Scan()
			input := scanner.Text()
			newBlock := blockchain.GenerateNewBlock(blockChain[len(blockChain) - 1], input)
			blockChain = append(blockChain, newBlock)
			fmt.Println("Block added!")
		case 2:
			blockchain.ViewBlockchain(blockChain)
		case 3:
			if blockchain.IsValidChain(blockChain) {
				fmt.Println("Blockchain is valid.")
			} else {
				fmt.Println("Blockchain integrity is compromised!")
			}
		case 4:
			blockchain.ResetBlockchain(&blockChain)
			fmt.Println("Blockchain reset to genesis block.")
		case 5:
			fmt.Println("Exiting program.")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
