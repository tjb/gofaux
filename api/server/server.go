package server

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gofaux/internal/writer"
	"log"
	"net/http"
	"os"
)

func CreateAndStartServer() {
	router := gin.Default()

	fileMap, err := writer.ParseFiles("tmp")
	if err != nil {
		log.Fatalf("Failed to parse files: %v\n", err)
	}

	for name, path := range fileMap {
		pathCopy := path // Avoid capturing loop variable
		router.GET(fmt.Sprintf("/%s", name), func(c *gin.Context) {
			serveJSON(c, pathCopy)
		})
	}

	err = router.Run(":6666")
	if err != nil {
		return
	}
}

func serveJSON(c *gin.Context, filePath string) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}

	var jsonData map[string]interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON"})
		return
	}

	c.JSON(http.StatusOK, jsonData)
}
