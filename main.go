package main

import (
	"fmt"

	"github.com/DmitryCodeGTS/builder-https-server/config"
)

func main() {
	fmt.Println("========== EXAMPLE 1: Default Configuration ==========\n")

	// Пример 1: Минимальная конфигурация (только значения по умолчанию)
	config1 := config.NewServerBuilder().Build()
	config1.Print()

	fmt.Println("\n========== EXAMPLE 2: Custom Configuration ==========\n")

	// Пример 2: Полная конфигурация с цепочкой вызовов
	config2 := config.NewServerBuilder().
		SetPort(9000).
		SetReadTimeout(30).
		SetWriteTimeout(30).
		EnableLogging().
		SetMaxHeaderBytes(2048576).
		Build()
	config2.Print()

	fmt.Println("\n========== EXAMPLE 3: Partial Configuration ==========\n")

	// Пример 3: Частичная конфигурация (некоторые значения по умолчанию)
	config3 := config.NewServerBuilder().
		SetPort(3000).
		EnableLogging().
		Build()
	config3.Print()
}
