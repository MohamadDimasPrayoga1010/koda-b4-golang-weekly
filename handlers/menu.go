package handlers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"main/utils"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func (m *Menu) InputMenu() {

	tempDir := os.TempDir()
	cacheFile := filepath.Join(tempDir, "data.json")

	var menuData []Menu

	info, err := os.Stat(cacheFile)
	if os.IsNotExist(err) {
		fmt.Println("File cache tidak ada, fetching data...")
		resp, err := http.Get("https://raw.githubusercontent.com/MohamadDimasPrayoga1010/koda-b4-golang-weekly-data/refs/heads/main/data.json")
		if err != nil {
			fmt.Println("Failed fetch data:", err)
			return
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		json.Unmarshal(body, &menuData)
		os.WriteFile(cacheFile, body, 0644)
	} else if err != nil {
		fmt.Println("Error cek cache:", err)
		return
	} else {
		age := time.Since(info.ModTime())
		if age >= 15*time.Minute {
			fmt.Println("Cache expired, fetching data baru...")
			resp, err := http.Get("https://raw.githubusercontent.com/MohamadDimasPrayoga1010/koda-b4-golang-weekly-data/refs/heads/main/data.json")
			if err != nil {
				fmt.Println("Failed fetch data:", err)
				return
			}
			defer resp.Body.Close()

			body, _ := io.ReadAll(resp.Body)
			json.Unmarshal(body, &menuData)
			os.WriteFile(cacheFile, body, 0644)
		} else {
			fmt.Println("Mengambil data dari cache")
			body, err := os.ReadFile(cacheFile)
			if err != nil {
				fmt.Println("Gagal baca cache, fetching data baru...")
				resp, err := http.Get("https://raw.githubusercontent.com/MohamadDimasPrayoga1010/koda-b4-golang-weekly-data/refs/heads/main/data.json")
				if err != nil {
					fmt.Println("Failed fetch data:", err)
					return
				}
				defer resp.Body.Close()

				body, _ := io.ReadAll(resp.Body)
				json.Unmarshal(body, &menuData)
				os.WriteFile(cacheFile, body, 0644)
			} else {
				json.Unmarshal(body, &menuData)
			}
		}
	}

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

		fmt.Printf("How many %s what you want to buy : ", selectedMenu.GetName())
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
