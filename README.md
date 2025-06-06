# Currency Converter CLI

Простое консольное приложение на Go для получения курсов валют и конвертации сумм между выбранными валютами с использованием API с ключом доступа.

---

## Возможности

- Получение актуального курса валют по базовой и целевой валюте
- Конвертация указанной суммы из одной валюты в другую
- Интерактивный ввод параметров через консоль
- Работа через API с обязательным API ключом

---

## Быстрый старт

1. Получи API ключ на [exchangerate.host](https://exchangerate.host/) или аналогичном сервисе.
2. Склонируй репозиторий или скачай исходники.
3. Вставь свой API ключ в переменную `accessKey` в файле `main.go`.
4. Запусти приложение:

```bash
go run main.go
```

5. Следуй подсказкам:

```
Введите базовую валюту (например, USD): USD  
Введите валюту для конвертации (например, EUR): EUR  
Введите сумму для конвертации: 100  
Результат: 100 USD = 87.95 EUR
```

---

## Поддерживаемые валюты (ISO 4217)

| Валюта               | Код  |
|----------------------|-------|
| Доллар США           | USD   |
| Евро                 | EUR   |
| Российский рубль     | RUB   |
| Британский фунт      | GBP   |
| Японская иена        | JPY   |
| Китайский юань       | CNY   |
| Швейцарский франк    | CHF   |
| Канадский доллар     | CAD   |
| Австралийский доллар | AUD   |
| Новозеландский доллар| NZD   |
| Индийская рупия      | INR   |
| Бразильский реал     | BRL   |
| Южноафриканский ранд | ZAR   |
| Мексиканский песо    | MXN   |
| Турецкая лира        | TRY   |

---

## Важные нюансы

- Для работы необходим действующий API ключ, вставленный в код.
- Валюты вводятся в формате ISO (например, USD, EUR).
- При ошибке запроса или неправильном коде валюты программа завершится с сообщением об ошибке.
- Требуется подключение к интернету.
- При бесплатном тарифе возможно 100 запросов в месяц

---
