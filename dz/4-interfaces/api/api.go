package api

import (
	"fmt"
	"struct/config"
)

func StartAPI() {
	config := &config.Config{}
	key := config.GetKey()
	fmt.Printf("Ключ получен: %v \n", key)
}
