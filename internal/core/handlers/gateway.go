package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"server/internal/core/models"
	"server/pkg/utils"

	"github.com/gin-gonic/gin"
)

// Mini-App Registry
var miniAppRegistry = make(map[string][]string)

type GatewayHandler struct{}

func NewGatewayHandler(r *gin.Engine) *GatewayHandler {
	handler := &GatewayHandler{}

	r.POST("/v1/gateway/register", handler.Register)
	r.POST("/v1/gateway/call", handler.Call)

	return handler
}

func (h *GatewayHandler) Register(c *gin.Context) {
	var data models.Resigter

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid registration data"})
		return
	}

	miniAppRegistry[data.AppName] = data.Functions
	log.Printf("📌 Registered Mini-App: %s\n", data.AppName)

	c.JSON(http.StatusOK, gin.H{"message": "Mini-App registered successfully!"})
}

func (h *GatewayHandler) Call(c *gin.Context) {
	var req models.Call

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	functions, exists := miniAppRegistry[req.TargetApp]
	if !exists || !utils.Contain(functions, req.FunctionName) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Function not found"})
		return
	}

	payloadBytes, err := json.Marshal(req.Payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error encoding payload"})
		return
	}

	url := fmt.Sprintf("http://localhost:3001/%s/%s", req.TargetApp, req.FunctionName)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error forwarding request"})
		return
	}
	defer response.Body.Close()

	var result map[string]interface{}
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading response"})
		return
	}

	if err := json.Unmarshal(responseBody, &result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing response"})
		return
	}

	c.JSON(http.StatusOK, result)
}
