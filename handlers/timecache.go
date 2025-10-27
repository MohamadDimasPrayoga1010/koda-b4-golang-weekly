package handlers

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func GetCache() time.Duration {

	cacheDuration := 15 * time.Second

	_ = godotenv.Load()

	if val := os.Getenv("CACHE_TIME_SECOND"); val != "" {
		if seconds, err := strconv.Atoi(val); err == nil && seconds > 0 {
			cacheDuration = time.Duration(seconds) * time.Second
		}
	}

	fmt.Printf("Cache TIME: %v\n", cacheDuration)
	return cacheDuration
}
