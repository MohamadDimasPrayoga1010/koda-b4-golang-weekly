package handlers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"main/utils"
)

func HistoryOrder() {
	defer func(){
		if r := recover(); r != nil{
			fmt.Println("Error, Program error captured:", r)
			fmt.Println("But don’t worry, returning to the menu...")
		}
	}()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("\x1bc")
		if len(History) == 0 {
			fmt.Println("\n==== HISTORY SHOPPING ====")
			fmt.Println("No shopping history")
		} else {
			fmt.Println("\n===============================")
			fmt.Println("====== HISTORY SHOPPING =======")
			fmt.Println("===============================")

			for i, invoice := range History {
				fmt.Printf("%d. Invoice: %s\n", i+1, invoice.InvoiceNumber)
				fmt.Printf("Date: %s\n", invoice.Date.Format("02/01/2006 15:04"))
				fmt.Println("-------------------------------")

				for j, order := range invoice.Orders {
					subtotal := order.Item.Price * order.Quantity
					fmt.Printf("%d. %s\n   Price: %s\n   Amount: %d\n   Subtotal: %s\n\n",
						j+1, order.Item.Name,
						utils.FormatInt64ToRp(int64(order.Item.Price)),
						order.Quantity,
						utils.FormatInt64ToRp(int64(subtotal)))
				}

				fmt.Printf("TOTAL SHOPPING: %s\n", utils.FormatInt64ToRp(int64(invoice.Total)))
				fmt.Println("===============================")
			}
		}

		fmt.Print("\nPress 0 to return: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "0" {
			fmt.Println("Return to menu")
			return
		} else {
			utils.SafePanic("Invalid input", reader)
		}
	}
}
