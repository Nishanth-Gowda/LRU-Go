package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nishanth-gowda/lru-golang/LRU"
)

func main() {
	router := gin.Default()

	cache := LRU.NewLRUCache(10)

	router.GET("/get", func(c *gin.Context) {
		key := c.Query("key")
		if key == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing key parameter"})
			return
		}

		value, ok := cache.Get(key)
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "key not found"})
			return
		}
		c.JSON(http.StatusOK, value)
	})

	router.POST("/set", func(c *gin.Context) {
		var data map[string]interface{}
		err := c.BindJSON(&data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
			return
		}

		key, ok := data["key"]
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing key in request body"})
			return
		}

		value, ok := data["value"]
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing value in request body"})
			return
		}

		var expiration time.Duration
		if exp, ok := data["expiration"]; ok {
			switch v := exp.(type) {
			case float64:
				expiration = time.Duration(v) * time.Second
			case int:
				expiration = time.Duration(v) * time.Second
			default:
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid expiration format"})
				return
			}
		}

		cache.Put(key, value, expiration)
		c.JSON(http.StatusOK, gin.H{"message": "key-value pair set successfully"})

	})

	router.Run(":8080")

}
