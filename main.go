package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/torbenconto/plutus/historical"
	"github.com/torbenconto/plutus/interval"
	prange "github.com/torbenconto/plutus/range"
	"github.com/torbenconto/plutus/stock"
	"net/http"
	"os"
)

func setupRouter() *gin.Engine {
	// Initialize gin router
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Authorization, Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})

	r.Use(cors.Default())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/quote/:ticker", func(c *gin.Context) {
		// Get ticker from url param
		ticker := c.Param("ticker")
		// Create new quote instance
		data, err := stock.NewQuote(ticker)
		// Check for errors, return 404 if not found or 200 along with quote data if found
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err,
			})
		} else {
			c.JSON(http.StatusOK, data)
		}

	})

	r.GET("/historical/:ticker", func(c *gin.Context) {
		// Get ticker from url param
		ticker := c.Param("ticker")
		// Get range from url param
		_range := prange.RangeFromString(c.Query("range"))
		// Get interval from url param
		_interval := interval.IntervalFromString(c.Query("interval"))
		// Create new historical instance
		data, err := historical.NewHistorical(ticker, _range, _interval)
		// Check for errors, return 404 if not found or 200 along with historical data if found
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err,
			})
		} else {
			c.JSON(http.StatusOK, data)
		}
	})

	r.GET("/dividend/:ticker", func(c *gin.Context) {
		ticker := c.Param("ticker")
		fmt.Printf("Received request for ticker: %s\n", ticker) // Log the ticker

		data, err := stock.NewDividendInfo(ticker)
		if err != nil {
			fmt.Printf("Error fetching dividend info for ticker %s: %v\n", ticker, err) // Log the error
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(), // Ensure the error is string
			})
		} else {
			c.JSON(http.StatusOK, data)
		}
	})

	return r
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8001" // Default port
	}

	// Setup router
	r := setupRouter()

	// Listen and serve on the specified port
	err := r.Run(":" + port)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
