package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Mini-App Registry
var miniAppRegistry = make(map[string][]string)

func NewSuperHandler(r *gin.Engine) {
	// Configure routes to match what the SDK expects
	r.GET("/v1/super/list", getlistMiniApp)
	r.POST("/v1/super/register", registerMiniApp)
	r.POST("/v1/super/call-function", callMiniAppFunction)
}

func registerMiniApp(c *gin.Context) {
	var data struct {
		AppName   string   `json:"appName"`
		Functions []string `json:"functions"`
	}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid registration data"})
		return
	}
	miniAppRegistry[data.AppName] = data.Functions
	log.Printf("📌 Registered Mini-App: %s\n", data.AppName)
	log.Printf("📌 Functions: %v\n", miniAppRegistry)
	c.JSON(http.StatusOK, gin.H{"message": "Mini-App registered successfully!"})
}

func getlistMiniApp(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"miniApps": miniAppRegistry,
	})
}

func callMiniAppFunction(c *gin.Context) {
	var req struct {
		Url          string         `json:"url"`
		Caller       string         `json:"caller"`
		TargetApp    string         `json:"targetApp"`
		FunctionName string         `json:"functionName"`
		Payload      map[string]any `json:"payload"`
	}

	// Read the raw request body
	requestData, _ := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestData))
	log.Printf("📥 Raw Request Body: %s\n", string(requestData))

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("❌ JSON Parsing Error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	log.Printf("✅ Parsed Request: %+v\n", req)

	functions, exists := miniAppRegistry[req.TargetApp]
	if !exists || !contains(functions, req.FunctionName) {
		log.Printf("❌ Function not found: %s.%s\n", req.TargetApp, req.FunctionName)
		c.JSON(http.StatusNotFound, gin.H{"error": "Function not found"})
		return
	}

	payloadBytes, err := json.Marshal(req.Payload)
	if err != nil {
		log.Printf("❌ Error encoding payload: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error encoding payload"})
		return
	}

	// Try multiple possible hostnames to connect to mini-app-b
	urls := []string{
		// use this URL if you are running the server locally
		fmt.Sprintf("%s/%s", req.Url, req.FunctionName),

		// use this URL if you are running the server in a Docker container
		fmt.Sprintf("http://localhost:8081/%s", req.FunctionName),
		fmt.Sprintf("http://host.docker.internal:8081/%s", req.FunctionName),
	}

	var responseBody []byte
	var responseErr error
	var successful bool

	// Try each URL until one works
	for _, url := range urls {
		log.Printf("🔄 Trying to forward request to: %s with payload: %s\n", url, string(payloadBytes))

		client := &http.Client{
			Timeout: 5 * time.Second,
		}

		request, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
		if err != nil {
			log.Printf("❌ Error creating request: %v\n", err)
			continue
		}

		request.Header.Set("Content-Type", "application/json")

		response, err := client.Do(request)
		if err != nil {
			log.Printf("❌ Error forwarding to %s: %v\n", url, err)
			responseErr = err
			continue
		}

		// Read the response body
		responseBody, err = ioutil.ReadAll(response.Body)
		response.Body.Close()

		if err != nil {
			log.Printf("❌ Error reading response from %s: %v\n", url, err)
			responseErr = err
			continue
		}

		log.Printf("✅ Successfully received response from %s: %s\n", url, string(responseBody))
		successful = true
		break
	}

	if !successful {
		log.Printf("❌ All connection attempts failed. Last error: %v\n", responseErr)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error forwarding request: %v", responseErr)})
		return
	}

	// Parse the response
	var result map[string]interface{}
	if err := json.Unmarshal(responseBody, &result); err != nil {
		log.Printf("❌ Error parsing response: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing response"})
		return
	}

	c.JSON(http.StatusOK, result)
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
