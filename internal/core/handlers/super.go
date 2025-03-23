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

// MiniAppInfo stores information about registered mini-apps
type MiniAppInfo struct {
	Functions []string
	URL       string
}

// Mini-App Registry with URL information
var miniAppRegistry = make(map[string]MiniAppInfo)

func NewSuperHandler(r *gin.Engine) {
	// Configure routes to match what the SDK expects
	r.GET("/v1/super/list", getlistMiniApp)
	r.POST("/v1/super/register", registerMiniApp)
	r.POST("/v1/super/call-function", callMiniAppFunction)
}

func registerMiniApp(c *gin.Context) {
	// Updated registration data structure to include URL
	var data struct {
		AppName   string   `json:"appName"`
		Functions []string `json:"functions"`
		URL       string   `json:"url"`
	}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid registration data"})
		return
	}

	// Make sure the URL doesn't end with a slash
	if len(data.URL) > 0 && data.URL[len(data.URL)-1] == '/' {
		data.URL = data.URL[:len(data.URL)-1]
	}

	// Store the mini-app information including URL
	miniAppRegistry[data.AppName] = MiniAppInfo{
		Functions: data.Functions,
		URL:       data.URL,
	}

	log.Printf("Registered Mini-App: %s\n", data.AppName)
	log.Printf("Functions: %v\n", data.Functions)
	log.Printf("URL: %s\n", data.URL)

	c.JSON(http.StatusOK, gin.H{"message": "Mini-App registered successfully!"})
}

func getlistMiniApp(c *gin.Context) {
	// Format the registry to include functions and URLs
	response := make(map[string]map[string]interface{})

	for appName, info := range miniAppRegistry {
		response[appName] = map[string]interface{}{
			"functions": info.Functions,
			"url":       info.URL,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"miniApps": response,
	})
}

func callMiniAppFunction(c *gin.Context) {
	var req struct {
		Caller       string         `json:"caller"`
		TargetApp    string         `json:"targetApp"`
		FunctionName string         `json:"functionName"`
		Payload      map[string]any `json:"payload"`
	}

	// Read the raw request body
	requestData, _ := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestData))
	log.Printf("Raw Request Body: %s\n", string(requestData))

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("JSON Parsing Error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	log.Printf("Parsed Request: %+v\n", req)

	// Look up the target mini-app info
	targetAppInfo, exists := miniAppRegistry[req.TargetApp]
	if !exists || !contains(targetAppInfo.Functions, req.FunctionName) {
		log.Printf("Function not found: %s.%s\n", req.TargetApp, req.FunctionName)
		c.JSON(http.StatusNotFound, gin.H{"error": "Function not found"})
		return
	}

	payloadBytes, err := json.Marshal(req.Payload)
	if err != nil {
		log.Printf("Error encoding payload: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error encoding payload"})
		return
	}

	// Get the base URL of the target Mini-App from the registry
	miniAppBaseURL := targetAppInfo.URL
	fmt.Println("asdfasdf", miniAppBaseURL)

	// Make sure the base URL doesn't end with a slash
	if len(miniAppBaseURL) > 0 && miniAppBaseURL[len(miniAppBaseURL)-1] == '/' {
		miniAppBaseURL = miniAppBaseURL[:len(miniAppBaseURL)-1]
	}

	// Construct the full URL using the convention: baseURL/v1/super/functionName
	url := fmt.Sprintf("%s/%s", miniAppBaseURL, req.FunctionName)
	log.Printf("Forwarding request to target app URL: %s with payload: %s\n", url, string(payloadBytes))

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Printf("Error creating request: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error creating request: %v", err)})
		return
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		log.Printf("Error forwarding to %s: %v\n", url, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error forwarding request: %v", err)})
		return
	}

	// Read the response body
	responseBody, err := ioutil.ReadAll(response.Body)
	response.Body.Close()

	if err != nil {
		log.Printf("Error reading response from %s: %v\n", url, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error reading response: %v", err)})
		return
	}

	log.Printf("Successfully received response from %s: %s\n", url, string(responseBody))

	// Parse the response
	var result map[string]interface{}
	if err := json.Unmarshal(responseBody, &result); err != nil {
		log.Printf("Error parsing response: %v\n", err)
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
