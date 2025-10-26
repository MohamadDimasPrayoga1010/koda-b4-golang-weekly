package handlers

import (
	"bufio"
	"fmt"
	"main/utils"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func Checkout() {
	defer func() {
		if r := recover(); r != nil {
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
				subtotal := order.GetSubtotal()
				fmt.Printf("%d. %s\n   Price: %s\n   Amount: %d\n   Subtotal: %s\n\n",
					i+1,
					order.GetItemName(),
					utils.FormatInt64ToRp(int64(order.GetItemPrice())),
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
				fmt.Println("The cart is empty, nothing to pay for!")
				break
			}

			PaymentShopping(reader)

			return

		case "0":
			fmt.Println("↩Returning to menu...")
			return

		default:
			fmt.Println("Invalid input!")
			time.Sleep(1 * time.Second)
		}
	}
}

func PaymentShopping(reader *bufio.Reader) {

	paymentMethod := []string{"Ovo", "Dana", "ShoppePay", "Bca", "Mandiri", "Bni", "Bri"}

	var wg sync.WaitGroup
	result := make(chan string, len(paymentMethod))

	fmt.Printf("\nLoading...Checking available payment methods...\n")

	for _, method := range paymentMethod {
		wg.Add(1)
		go func(method string) {
			defer wg.Done()

			rand.Seed(time.Now().UnixNano())
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)

			result <- method
		}(method)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	var available []string
	for res := range result {
		available = append(available, res)
	}

	fmt.Println("Payment methods available:")
	for i, p := range available {
		fmt.Printf("%d. %s\n", i+1, p)
	}

	var selectedPayment string
	for {
		fmt.Print("\nSelect payment method: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		choice, err := strconv.Atoi(input)
		if err != nil || choice < 1 || choice > len(available) {
			fmt.Println("\nInvalid choice! Please try again.")
			continue 
		}

		selectedPayment = strings.Split(available[choice-1], " ")[0]
		break 
	}

	fmt.Printf("\nProcessing your payment with %s...\n", selectedPayment)
	rand.Seed(time.Now().UnixNano())
	paymentTime := rand.Intn(5)
	time.Sleep(time.Duration(paymentTime) * time.Second)

	invoice := Invoice{
		InvoiceNumber: utils.RandomInvoce(10),
		Date:          time.Now(),
		Orders:        Orders,
		Total:         0,
	}

	invoice.Total = invoice.CalculateTotal()

	fmt.Printf("\nPayment with %s successful in %d seconds!\n", selectedPayment, paymentTime)
	fmt.Printf("Invoice Number : %s\n", invoice.InvoiceNumber)
	fmt.Printf("Date           : %s\n", invoice.Date.Format("02/01/2006 15:04"))
	fmt.Printf("Total Payment  : %s\n", utils.FormatInt64ToRp(int64(invoice.Total)))
	fmt.Printf("Items Purchased: %d\n", invoice.GetOrderCount())

	History = append(History, invoice)
	Orders = nil

	fmt.Println("\nPress Enter to return...")
	reader.ReadString('\n')
}
