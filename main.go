package main

import (
	"fmt"
	"os"

	"e-wallet-api/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.GetConn()

	router := gin.Default()

	version := os.Getenv("API_VERSION")
	router.Group(fmt.Sprintf("/api/%s", version))

	router.Run(":8000")
}
