package main

import (
	"fmt"
	"os"
)

// Выводит в консоль переменные окружения и настраивает по умолчанию, если они не заданы.
func env() {
	fmt.Println("Enviroment variables:")

	if os.Getenv("FAKE_API_PORT") == "" {
		os.Setenv("FAKE_API_PORT", "8080")
	}

	fmt.Println("FAKE_API_PORT:", os.Getenv("FAKE_API_PORT"))

	if os.Getenv("FAKE_API_HOST") == "" {
		os.Setenv("FAKE_API_HOST", "0.0.0.0")
	}

	fmt.Println("FAKE_API_HOST:", os.Getenv("FAKE_API_HOST"))

	if os.Getenv("FAKE_API_MAX_FIELDS_IN_OBJECT") == "" {
		os.Setenv("FAKE_API_MAX_FIELDS_IN_OBJECT", "50")
	}

	fmt.Println("FAKE_API_MAX_FIELDS_IN_OBJECT:", os.Getenv("FAKE_API_MAX_FIELDS_IN_OBJECT"))
}

// Главная точка входа в приложение. Запускает http сервер.
func main() {
	env()
	runHttpServer()
}
