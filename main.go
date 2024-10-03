package main

import (
    "github.com/dev-geov/crypto-api/controllers"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    // Definição das rotas
    router.GET("/cryptos", controllers.ListCryptos)
    router.GET("/crypto/:id", controllers.GetCryptoInfo)
    router.POST("/convert", controllers.ConvertCrypto)

    router.Run(":8080") // Inicia o servidor na porta 8080
}
