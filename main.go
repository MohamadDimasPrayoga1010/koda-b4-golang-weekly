package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"main/handlers"
)

func main() {
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
			fmt.Println("Input harus berupa angka!")
			fmt.Println("Tekan Enter untuk mengisi ulang...")
			reader.ReadString('\n')
			continue
		}

		switch choose {
		case 1:
			menu.InputMenu()
		case 2:
			handlers.Checkout()
		case 3:
			handlers.HistoryOrder()
		case 4:
			fmt.Println("Terima kasih sudah belanja!")
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak valid!")
			fmt.Println("Tekan Enter untuk mengisi ulang...")
			reader.ReadString('\n')
		}
	}
}
