package handlers

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Option() {

	tempDir := filepath.Join(os.TempDir(), "burgerbangor")
	cacheFile := filepath.Join(tempDir, "data.json")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("\x1bc")
		fmt.Println("============OPTION============")
		fmt.Println("\n1. Clear cache")
		fmt.Println("\n0. Back to main menu")
		fmt.Print("\nChoise : ")

		var choise int
		fmt.Scanln(&choise)

		// input, _ := reader.ReadString('\n')
		// input = strings.TrimSpace(input)

		switch choise {
		case 1:
			if _, err := os.Stat(cacheFile); os.IsNotExist(err) {
				fmt.Println("Cache file tidak ditemukan.")
				break
			}

			fmt.Print("Apakah Anda yakin ingin menghapus cache? (y/n): ")
			confirm, _ := reader.ReadString('\n')
			confirm = strings.TrimSpace(strings.ToLower(confirm))

			if confirm == "y" {
				if err := os.Remove(cacheFile); err != nil {
					fmt.Println("Gagal menghapus cache:", err)
				} else {
					fmt.Println("Cache berhasil dihapus!")
				}
			} else {
				fmt.Println("Invalid input, Hapus cache dibatalkan.")
			}

		case 0:
			return

		default:
		fmt.Println("Input Tidak Valid")
		}
		fmt.Println("\nTekan Enter untuk kembali ke Option Menu...")
		reader.ReadString('\n')
	}
}
