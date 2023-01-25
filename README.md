# Golang scraper example
Небольшой скрипт для скрапинга сайта [books.toscrape.com](https://books.toscrape.com/) (пакет `goquery`)
- сбор информации о книгах: наименование, ссылка на страницу книги, цена, наличие;
- запись полученной информации в `csv-файл`.

## Requirements
Go 1.19 or above.

## Installation
```
git clone https://github.com/s-antoshkin/go-scraper-example
cd go-scraper-example
go get -u
```

## Start the application
```
go run .
```

## Packages used
```
"encoding/csv"
"fmt"
"net/http"
"os"
"regexp"
"strings"
"time"

"github.com/PuerkitoBio/goquery"
```