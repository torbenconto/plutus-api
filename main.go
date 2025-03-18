package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/torbenconto/plutus/v2"
	"github.com/torbenconto/plutus/v2/interval"
	_range "github.com/torbenconto/plutus/v2/range"
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
		data, err := plutus.GetQuote(ticker)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, data)
		}

	})

	r.GET("/historical/:ticker", func(c *gin.Context) {
		// Get ticker from url param
		ticker := c.Param("ticker")
		// Get range from url param
		r := _range.RangeFromString(c.Query("range"))
		// Get interval from url param
		i := interval.IntervalFromString(c.Query("interval"))
		// Create new historical instance

		data, err := plutus.GetHistoricalQuote(ticker, r, i)
		// Check for errors, return 404 if not found or 200 along with historical data if found
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, data)
		}
	})

	r.GET("/dividend/:ticker", func(c *gin.Context) {
		ticker := c.Param("ticker")
		fmt.Printf("Received request for ticker: %s\n", ticker) // Log the ticker

		data, err := plutus.GetDividendInfo(ticker)
		if err != nil {
			fmt.Printf("Error fetching dividend info for ticker %s: %v\n", ticker, err) // Log the error
			c.JSON(http.StatusInternalServerError, gin.H{
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
		port = "8080" // Default port
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
