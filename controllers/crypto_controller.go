package controllers

import (
    "github.com/dev-geov/crypto-api/services"
    "github.com/gin-gonic/gin"
    "net/http"
)

// Rota para listar as criptomoedas e seus preços
func ListCryptos(c *gin.Context) {
    cryptos, err := services.GetCryptoPrices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error retrieving crypto prices"})
		return
	}
	c.JSON(http.StatusOK, cryptos)

}

// Rota para converter criptomoeda para a moeda escolhida
func ConvertCrypto(c *gin.Context) {
    var conversionRequest services.ConversionRequest
    if err := c.BindJSON(&conversionRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Dados de entrada inválidos"})
        return
    }
    result, err := services.ConvertCrypto(conversionRequest)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, result)
}

// Rota para obter informações sobre uma criptomoeda específica
func GetCryptoInfo(c *gin.Context) {
    id := c.Param("id")
    info, err := services.GetCryptoInfo(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Criptomoeda não encontrada"})
        return
    }
    c.JSON(http.StatusOK, info)
}
