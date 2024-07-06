package main

import (
	"encoding/json"
	"github.com/google/uuid"
	_ "halyk-life_proxy_server/docs"
	"halyk-life_proxy_server/pkg/proxy/models"
	"io"
	"net/http"
	"sync"
)

type App struct {
}

var storage sync.Map

// handleProxy godoc
// @Summary Handles the proxy request
// @Description This endpoint proxies a request.
// @Tags proxy
// @Accept  json
// @Produce  json
// @Param request body models.ProxyRequest true "Proxy request"
// @Success 200 {object} models.ProxyResponse
// @Router /proxy [post]
func (a *App) handleProxy(w http.ResponseWriter, r *http.Request) {
	var req models.ProxyRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	client := &http.Client{}
	httpReq, err := http.NewRequest(req.Method, req.URL, nil)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	for key, value := range req.Headers {
		httpReq.Header.Set(key, value)
	}

	resp, err := client.Do(httpReq)
	if err != nil {
		http.Error(w, "Failed to proxy request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	headers := make(map[string]string)
	for key, values := range resp.Header {
		headers[key] = values[0]
	}

	proxyResp := models.ProxyResponse{
		ID:      uuid.New().String(),
		Status:  resp.StatusCode,
		Headers: headers,
		Length:  len(body),
	}

	storage.Store(proxyResp.ID, proxyResp)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(proxyResp)
}

// retrieveHistory godoc
// @Summary Retrieves request history
// @Description This endpoint retrieves the history of proxied requests.
// @Tags history
// @Accept  json
// @Produce  json
// @Success 200 {array} models.ProxyResponse
// @Router /history [get]
func (a *App) retrieveHistory(w http.ResponseWriter, r *http.Request) {
	var history []models.ProxyResponse
	storage.Range(func(key, value interface{}) bool {
		history = append(history, value.(models.ProxyResponse))
		return true
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(history)
}

// healthCheck godoc
// @Summary Health check
// @Description This endpoint checks the health of the server.
// @Tags health
// @Accept  json
// @Produce  json
// @Success 200 {string} string "OK"
// @Router /health [get]
func (a *App) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
