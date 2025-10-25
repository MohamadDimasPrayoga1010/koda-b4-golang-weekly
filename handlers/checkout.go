package handlers

import (
	"bufio"
	"fmt"
	"main/utils"
	"os"
	"strings"
	"time"
)

func Checkout() {
	defer func(){
		if r := recover(); r != nil{
			fmt.Println("Error, Program error captured:", r)
			fmt.Println("But don’t worry, returning to the menu...")
		}
	}()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("\x1bc")
		if len(Orders) == 0 {
			fmt.Println("\nEmpty shopping cart!")
		} else {
			fmt.Println("\n=== SHOPPING CART ===")
			total := 0
			for i, order := range Orders {
				subtotal := order.Item.Price * order.Quantity
				fmt.Printf("%d. %s\n   Price: %s\n   Amount: %d\n   Subtotal: %s\n\n",
					i+1,
					order.Item.Name,
					utils.FormatInt64ToRp(int64(order.Item.Price)),
					order.Quantity,
					utils.FormatInt64ToRp(int64(subtotal)),
				)
				total += subtotal
			}
			fmt.Println("===========================")
			fmt.Printf("TOTAL SHOPPING: %s\n", utils.FormatInt64ToRp(int64(total)))
			fmt.Println("===========================")
		}

		fmt.Print("\nOrder checkout (y/0 to return): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		switch input {
		case "y":
			if len(Orders) == 0 {
				fmt.Println("The cart is empty, nothing to pay for")
			} else {
				total := 0
				for _, order := range Orders {
					total += order.Item.Price * order.Quantity
				}

				invoice := Invoice{
					InvoiceNumber: utils.RandomInvoce(10),
					Date:          time.Now(),
					Orders:        Orders,
					Total:         total,
				}

				History = append(History, invoice)
				Orders = nil

				fmt.Println("\nPayment successful!")
				fmt.Printf("Invoice: %s\n", invoice.InvoiceNumber)
				fmt.Printf("Date: %s\n", invoice.Date.Format("02/01/2006 15:04"))
				fmt.Printf("Total payment: %s\n", utils.FormatInt64ToRp(int64(invoice.Total)))
			}
			fmt.Println("\nPress Enter to return...")
			reader.ReadString('\n')
			return

		case "0":
			fmt.Println("Return to menu")
			return

		default:
			utils.SafePanic("Invalid input", reader)
		}
	}
}
