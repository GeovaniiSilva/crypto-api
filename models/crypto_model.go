package models

type Crypto struct {
    ID     string  `json:"id"`
    Name   string  `json:"name"`
    Price  float64 `json:"price"`
    Symbol string  `json:"symbol"`
}
