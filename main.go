package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type RatesResponse struct {
	Success bool               `json:"success"`
	Source  string             `json:"source"`
	Quotes  map[string]float64 `json:"quotes"`
}

func main() {
	var amount float64
	var base, target string

	fmt.Print("Введите сумму для конвертации: ")
	_, err := fmt.Scanln(&amount)
	if err != nil {
		log.Fatalf("Ошибка ввода суммы: %v", err)
	}

	fmt.Print("Введите валюту, из которой переводим (например, USD): ")
	_, err = fmt.Scanln(&base)
	if err != nil {
		log.Fatalf("Ошибка ввода базовой валюты: %v", err)
	}

	fmt.Print("Введите валюту, в которую переводим (например, EUR): ")
	_, err = fmt.Scanln(&target)
	if err != nil {
		log.Fatalf("Ошибка ввода целевой валюты: %v", err)
	}

	base = strings.ToUpper(base)
	target = strings.ToUpper(target)

	url := fmt.Sprintf("http://api.currencylayer.com/live?access_key=0cc2336080c97552dff5403944321310&source=%s&currencies=%s", base, target)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Ошибка запроса: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Ошибка HTTP: %v", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Ошибка чтения тела ответа: %v", err)
	}

	var rates RatesResponse
	if err := json.Unmarshal(body, &rates); err != nil {
		log.Fatalf("Ошибка парсинга JSON: %v", err)
	}

	if !rates.Success {
		log.Fatalf("API вернул неуспешный ответ")
	}

	key := rates.Source + target
	rate, ok := rates.Quotes[key]
	if !ok {
		log.Fatalf("Валюта %s не найдена в ответе", target)
	}

	converted := amount * rate
	fmt.Printf("%.2f %s = %.4f %s\n", amount, base, converted, target)
}
