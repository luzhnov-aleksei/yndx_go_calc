# Web Calculator API

## Описание
Этот проект представляет собой веб-сервис, который вычисляет арифметические выражения. Пользователь отправляет арифметическое выражение через HTTP-запрос, а сервис возвращает результат вычисления.

## Основной endpoint

- **URL**: /api/v1/calculate
- **Метод**: POST
- **Тело запроса** (JSON):
  
```json
{
    "expression": "выражение, которое ввёл пользователь"
}
```
## Ответы от сервиса
Успешный ответ (HTTP 200): Если выражение успешно вычислено:

```json
{
  "result": "результат выражения"
}
```
Ошибка 422 (Неверное выражение): Если выражение невалидно (например, содержит недопустимые символы):

```json
{
  "error": "Expression is not valid"
}
```
Ошибка 500 (Внутренняя ошибка сервера): Если произошла неизвестная ошибка:

```json
{
  "error": "Internal server error"
}
```
## Установка и запуск
Клонируйте репозиторий на свою машину:

```bash
git clone https://github.com/yourusername/yndx_go_calc.git
```
Перейдите в каталог проекта:

```bash
cd yndx_go_calc
```
Для запуска сервиса используйте команду:

```bash
go run ./cmd/calc_service/...
```
По умолчанию сервис будет запущен на порту 8080.

Пример использования с curl
Пример успешного запроса: Запрос с выражением 2+2*2:

```bash
curl --location "http://localhost:8080/api/v1/calculate" ^
--header "Content-Type: application/json" ^
--data "{ \"expression\": \"2+2*2\" }"
```
Ответ:

```json
{
  "result": "6.000000"
}
```
Пример запроса с ошибкой 422: Запрос с неверным выражением, содержащим буквы:

```bash
curl --location "http://localhost:8080/api/v1/calculate" ^
--header "Content-Type: application/json" ^
--data "{ \"expression\": \"2+2a\" }"
```
Ответ:

```json
{
  "error": "Expression contains invalid characters"
}
```
Пример запроса с ошибкой 500: Запрос с некорректным JSON:

```bash
curl --location "http://localhost:8080/api/v1/calculate" ^
--header "Content-Type: application/json" ^
--data "{ "expression\": \"2+2*2\" }"
```
Ответ:

```json
{
  "error": "Internal server error"
}
```
## Тестирование
Для запуска тестов используйте команду:

```bash
go test ./...
```
#### Тесты
В проекте присутствуют следующие тесты:

Тесты на проверку выражений с недопустимыми символами.
Тесты для обработки запросов и ответов через HTTP.