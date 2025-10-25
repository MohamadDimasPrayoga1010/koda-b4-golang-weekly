package handlers

import (
	"bufio"
	"fmt"
	"main/utils"
	"os"
	"strconv"
	"strings"
)

var MenuData = []Menu{
	{ID: 1, Name: "Bangor Pitik Lava", Price: 29000},
	{ID: 2, Name: "Bangor Pitik Lava Premium", Price: 29000},
	{ID: 3, Name: "Bangor Cheese Lava", Price: 31000},
	{ID: 4, Name: "Bangor Lava Sausage", Price: 27500},
	{ID: 5, Name: "Bangor Jelata Cheese", Price: 24700},
	{ID: 6, Name: "Bangor Juragan Cheese", Price: 31700},
	{ID: 7, Name: "Bangor Ningrat Cheese", Price: 49200},
	{ID: 8, Name: "Bangor Juragan", Price: 29000},
	{ID: 9, Name: "Bangor Ningrat", Price: 44200},
	{ID: 10, Name: "Bangor Sultan", Price: 55500},
	{ID: 11, Name: "Bangor Fish", Price: 27500},
	{ID: 12, Name: "Tea", Price: 9500},
	{ID: 13, Name: "Soft Drink", Price: 10500},
}

func (m *Menu) InputMenu() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error, Program error captured:", r)
			fmt.Println("But don’t worry, returning to the menu...")
			m.InputMenu()
		}
	}()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("\x1bc")
		fmt.Println("\n=== Bangor Burger Menu List ===")
		for _, menu := range MenuData {
			fmt.Printf("%d. %s - %s\n",
				menu.ID,
				menu.Name,
				utils.FormatInt64ToRp(int64(menu.Price)),
			)
		}
		fmt.Println("0. Return to main menu")

		if len(Orders) > 0 {
			fmt.Println("\n===================== CURRENT ORDERS ===================== ")
			total := 0
			for i, o := range Orders {
				sub := o.Item.Price * o.Quantity
				fmt.Printf("%d. %s - %s x%d = %s\n",
					i+1,
					o.Item.Name,
					utils.FormatInt64ToRp(int64(o.Item.Price)),
					o.Quantity,
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
		for _, menu := range MenuData {
			if menu.ID == menuID {
				selectedMenu = &menu
				break
			}
		}

		if selectedMenu == nil {
			utils.SafePanic("The menu was not found", reader)
		}

		fmt.Printf("How many %s what you want to buy : ", selectedMenu.Name)
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

		subtotal := selectedMenu.Price * qty
		fmt.Println("\n============================================================")
		fmt.Printf("%d x %s added to the order!\nSubtotal: %s\n",
			qty,
			selectedMenu.Name,
			utils.FormatInt64ToRp(int64(subtotal)),
		)
		fmt.Println("============================================================")

		utils.Alert("\nPress Enter to return to the menu list...")
		reader.ReadString('\n')
	}
}
