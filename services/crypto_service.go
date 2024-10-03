package services

import (
    "encoding/json"
    "errors"
    "fmt"
    "net/http"
)

type Crypto struct {
    ID     string  `json:"id"`
    Name   string  `json:"name"`
    Price  float64 `json:"price"`
    Symbol string  `json:"symbol"`
}

type ConversionRequest struct {
    Amount       float64 `json:"amount"`
    CurrencyFrom string  `json:"currency_from"`
    CurrencyTo   string  `json:"currency_to"`
}

// Response do CoinGecko para preços de criptomoedas
type CoinGeckoPriceResponse map[string]map[string]float64

// Obter lista de criptomoedas e seus preços em BRL e USD usando a API CoinGecko
func GetCryptoPrices() ([]Crypto, error) {
    url := "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin,ethereum,ripple,cardano,solana&vs_currencies=brl,usd"
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var result CoinGeckoPriceResponse
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, err
    }

    cryptos := []Crypto{
        {ID: "bitcoin", Name: "Bitcoin", Price: result["bitcoin"]["brl"], Symbol: "BTC"},
        {ID: "ethereum", Name: "Ethereum", Price: result["ethereum"]["brl"], Symbol: "ETH"},
        {ID: "ripple", Name: "Ripple", Price: result["ripple"]["brl"], Symbol: "XRP"},
        {ID: "cardano", Name: "Cardano", Price: result["cardano"]["brl"], Symbol: "ADA"},
        {ID: "solana", Name: "Solana", Price: result["solana"]["brl"], Symbol: "SOL"},
    }
    return cryptos, nil
}

// Realiza a conversão entre criptomoedas usando CoinGecko
func ConvertCrypto(req ConversionRequest) (map[string]float64, error) {
    url := fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s,%s&vs_currencies=brl,usd", req.CurrencyFrom, req.CurrencyTo)
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var result CoinGeckoPriceResponse
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, err
    }

    fromPrice := result[req.CurrencyFrom]["brl"]
    toPrice := result[req.CurrencyTo]["brl"]
    if fromPrice == 0 || toPrice == 0 {
        return nil, errors.New("Preço da criptomoeda não encontrado")
    }

    return map[string]float64{req.CurrencyTo: (req.Amount * fromPrice) / toPrice}, nil
}

// Obter informações detalhadas sobre uma criptomoeda específica usando CoinGecko
func GetCryptoInfo(id string) (*Crypto, error) {
    url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/%s", id)
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var cryptoInfo struct {
        ID     string `json:"id"`
        Name   string `json:"name"`
        Symbol string `json:"symbol"`
        MarketData struct {
            CurrentPrice map[string]float64 `json:"current_price"`
        } `json:"market_data"`
    }
    if err := json.NewDecoder(resp.Body).Decode(&cryptoInfo); err != nil {
        return nil, err
    }

    crypto := &Crypto{
        ID:     cryptoInfo.ID,
        Name:   cryptoInfo.Name,
        Symbol: cryptoInfo.Symbol,
        Price:  cryptoInfo.MarketData.CurrentPrice["brl"],
    }

    return crypto, nil
}
