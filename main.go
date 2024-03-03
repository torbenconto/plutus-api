package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/torbenconto/plutus/historical"
	"github.com/torbenconto/plutus/interval"
	"github.com/torbenconto/plutus/news"
	"github.com/torbenconto/plutus/quote"
	prange "github.com/torbenconto/plutus/range"
	"net/http"
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
		stock, err := quote.NewQuote(ticker)
		// Check for errors, return 404 if not found or 200 along with quote data if found
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err,
			})
		} else {
			c.JSON(http.StatusOK, stock)
		}

	})

	r.GET("/historical/:ticker/:range/:interval", func(c *gin.Context) {
		// Get ticker from url param
		ticker := c.Param("ticker")
		// Get range from url param
		_range := prange.RangeFromString(c.Param("range"))
		// Get interval from url param
		_interval := interval.IntervalFromString(c.Param("interval"))
		// Create new historical instance
		stock, err := historical.NewHistorical(ticker, _range, _interval)
		// Check for errors, return 404 if not found or 200 along with historical data if found
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err,
			})
		} else {
			c.JSON(http.StatusOK, stock)
		}
	})

	r.GET("/news/:query", func(c *gin.Context) {
		// Get query from url param
		query := c.Param("query")
		// Create new news instance
		data, err := news.NewNews(query)
		// Check for errors, return 404 if not found or 200 along with news data if found
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err,
			})
		} else {
			c.JSON(http.StatusOK, data)
		}
	})

	return r
}

func main() {
	r := setupRouter()

	// Listen and serve on 8000
	err := r.Run(":8000")
	if err != nil {
		return
	}
}
