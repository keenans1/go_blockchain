package user_input

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
