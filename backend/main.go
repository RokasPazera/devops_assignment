package main

import (
    "backend/database"
    "backend/routes"
    "os"
    "github.com/gin-gonic/gin"
)

func main() {
    database.Connect()

    // Read environment variables or use default values
    port := os.Getenv("SERVER_PORT")
    if port == "" {
        port = "8000"
    }

    // Setup GIN router with CORS middleware
    r := gin.Default()
    r.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
    })
    routes.Register(r)

    // Start the server
    // FIX
    if err := r.Run(":" + port); err != nil {
        panic("failed to start server: " + err.Error())
    }

    // TODO: replace previous lines with this in student version
    // r.Run(":" + port)
}
