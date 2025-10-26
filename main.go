package main

import (
	"bufio"
	"fmt"
	"main/handlers"
	"main/utils"
	"os"
	"strconv"
	"strings"
)

func main() {
	defer func(){
		if r := recover(); r!= nil{
			fmt.Println("Error, Program error captured")
			fmt.Println("But don’t worry, returning to the menu...")
			main()
		}
	}()
	reader := bufio.NewReader(os.Stdin)
	menu := &handlers.Menu{}

	for {
		fmt.Println("===== Burger Bangor =====")
		fmt.Println("\n1. Menu")
		fmt.Println("2. Checkout")
		fmt.Println("3. History")
		fmt.Println("4. Exit")
		fmt.Print("\nPilih Input: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		choose, err := strconv.Atoi(input)
		if err != nil {
			utils.SafePanic("Invalid input! Must be a number", reader)
		}

		switch choose {
		case 1:
			menu.InputMenu()
		case 2:
			handlers.Checkout()
		case 3:
			handlers.HistoryOrder()
		case 4:
			fmt.Println("Thank you for shopping!")
			os.Exit(0)
		default:
			utils.SafePanic("Invalid input! select input 1 to 4", reader)
		}
	}
}
