package handlers

import (
	"bufio"
	"context"
	"fmt"
	"main/utils"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type CartItem struct {
	ID, ProductID, Price, Quantity int
	Name                           string
}

func Checkout() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error, Program error captured:", r)
			fmt.Println("But don’t worry, returning to the menu...")
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	_ = godotenv.Load()
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}
	defer conn.Close(context.Background())

	userID := 1

	for {
		fmt.Printf("\x1bc")
		rows, err := conn.Query(context.Background(),
			`SELECT cart.id, cart.product_id, products.name, products.price, cart.quantity
	 FROM cart
	 JOIN products ON products.id = cart.product_id
	 WHERE cart.user_id=$1`, userID)
		if err != nil {
			fmt.Println("Gagal mengambil data keranjang:", err)
			return
		}

		var cartItems []CartItem
		var total int

		for rows.Next() {
			var item CartItem
			if err := rows.Scan(&item.ID, &item.ProductID, &item.Name, &item.Price, &item.Quantity); err != nil {
				fmt.Println("Failed to scan cart item:", err)
				return
			}
			cartItems = append(cartItems, item)
			total += item.Price * item.Quantity
		}
		rows.Close()

		if len(cartItems) == 0 {
			fmt.Println("\nEmpty shopping cart!")
		} else {
			fmt.Println("\n=== SHOPPING CART ===")
			for i, item := range cartItems {
				fmt.Printf("%d. %s\n   Price: %s\n   Amount: %d\n   Subtotal: %s\n\n",
					i+1,
					item.Name,
					utils.FormatInt64ToRp(int64(item.Price)),
					item.Quantity,
					utils.FormatInt64ToRp(int64(item.Price*item.Quantity)),
				)
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
			if len(cartItems) == 0 {
				fmt.Println("The cart is empty, nothing to pay for!")
				time.Sleep(1 * time.Second)
				break
			}
			PaymentShopping(reader, conn, userID, cartItems, total)
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

func PaymentShopping(reader *bufio.Reader, conn *pgx.Conn, userID int, cartItems []CartItem, total int) {
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
		Orders:        []Order{},
		Total:         total,
	}

	_, err := conn.Exec(context.Background(),
		`INSERT INTO orders (invoice, user_id, total, created_at)
		 VALUES ($1, $2, $3, $4)`,
		invoice.InvoiceNumber, userID, invoice.Total, invoice.Date)
	if err != nil {
		fmt.Println("Failed to insert into orders:", err)
		return
	}

	var orderID int
	err = conn.QueryRow(context.Background(),
		`SELECT id FROM orders WHERE invoice = $1`, invoice.InvoiceNumber).Scan(&orderID)
	if err != nil {
		fmt.Println("Failed to get order id:", err)
		return
	}

	for _, item := range cartItems {
		subtotal := item.Price * item.Quantity
		_, err = conn.Exec(context.Background(),
			`INSERT INTO order_items (order_id, product_id, quantity, subtotal)
			 VALUES ($1, $2, $3, $4)`,
			orderID, item.ProductID, item.Quantity, subtotal)
		if err != nil {
			fmt.Println("Failed to insert order items:", err)
			return
		}
	}

	_, _ = conn.Exec(context.Background(), "DELETE FROM cart WHERE user_id=$1", userID)

	Orders = []Order{}

	fmt.Printf("\nPayment with %s successful in %d seconds!\n", selectedPayment, paymentTime)
	fmt.Printf("Invoice Number : %s\n", invoice.InvoiceNumber)
	fmt.Printf("Date           : %s\n", invoice.Date.Format("02/01/2006 15:04"))
	fmt.Printf("Total Payment  : %s\n", utils.FormatInt64ToRp(int64(invoice.Total)))
	fmt.Printf("Items Purchased: %d\n", len(cartItems))

	History = append(History, invoice)

	fmt.Println("\nPress Enter to return...")
	reader.ReadString('\n')
}
