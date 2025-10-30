package handlers

import (
	"bufio"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"main/utils"
	"os"
	"strconv"
	"strings"
)

func (m *Menu) InputMenu() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load .env")
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println("Failed to connect:", err)
		return
	}
	defer conn.Close(context.Background())

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error, Program error captured:", r)
			fmt.Println("But don’t worry, returning to the menu...")
			m.InputMenu()
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	userID := 1 
	for {
		rows, err := conn.Query(context.Background(),
			"SELECT id, name, price, created_at, updated_at FROM products ORDER BY id ASC",
		)
		if err != nil {
			fmt.Println("Failed to get data:", err)
			return
		}

		var menuData []Menu
		for rows.Next() {
			var p Menu
			err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.CreatedAt, &p.UpdatedAt)
			if err != nil {
				fmt.Println("Failed to scan data:", err)
				return
			}
			menuData = append(menuData, p)
		}
		rows.Close()

		fmt.Printf("\x1bc")
		fmt.Println("\n=== Bangor Burger Menu List ===")
		for _, menu := range menuData {
			fmt.Printf("%d. %s - %s\n",
				menu.GetID(),
				menu.GetName(),
				utils.FormatInt64ToRp(int64(menu.GetPrice())),
			)
		}
		fmt.Println("0. Return to main menu")

		if len(Orders) > 0 {
			fmt.Println("\n===================== CURRENT ORDERS ===================== ")
			total := 0
			for i, order := range Orders {
				sub := order.GetSubtotal()
				fmt.Printf("%d. %s - %s x%d = %s\n",
					i+1,
					order.GetItemName(),
					utils.FormatInt64ToRp(int64(order.GetItemPrice())),
					order.Quantity,
					utils.FormatInt64ToRp(int64(sub)),
				)
				total += sub
			}
			fmt.Printf("\nTEMPORARY TOTAL: %s\n", utils.FormatInt64ToRp(int64(total)))
			fmt.Println("================================================")
		}

		fmt.Print("\nSelect menu : ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		menuID, err := strconv.Atoi(input)
		if err != nil {
			utils.SafePanic("Invalid input! Must be a number", reader)
		}

		if menuID == 0 {
			return
		}

		var selectedMenu *Menu
		for _, menu := range menuData {
			if menu.GetID() == menuID {
				selectedMenu = &menu
				break
			}
		}

		if selectedMenu == nil {
			utils.SafePanic("The menu was not found", reader)
		}

		fmt.Printf("How many %s you want to buy : ", selectedMenu.GetName())
		qtyInput, _ := reader.ReadString('\n')
		qtyInput = strings.TrimSpace(qtyInput)
		qty, err := strconv.Atoi(qtyInput)
		if err != nil || qty <= 0 {
			utils.SafePanic("Invalid amount! Must be greater than 0", reader)
		}

		order := Order{
			Item:     *selectedMenu,
			Quantity: qty,
		}
		Orders = append(Orders, order)

		_, err = conn.Exec(context.Background(),
			`INSERT INTO cart (user_id, product_id, quantity) VALUES ($1, $2, $3)
			 ON CONFLICT (user_id, product_id) DO UPDATE SET quantity = cart.quantity + EXCLUDED.quantity`,
			userID, selectedMenu.GetID(), qty,
		)
		if err != nil {
			fmt.Println("Failed to insert/update cart:", err)
		}

		subtotal := order.GetSubtotal()
		fmt.Println("\n============================================================")
		fmt.Printf("%d x %s added to the order!\nSubtotal: %s\n",
			qty,
			selectedMenu.GetName(),
			utils.FormatInt64ToRp(int64(subtotal)),
		)
		fmt.Println("============================================================")

		utils.Alert("\nPress Enter to return to the menu list...")
		reader.ReadString('\n')
	}
}

