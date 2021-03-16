package gohelpers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

func postExecutionLog(start time.Time, message string, success int, logUrl string, apiKey string) (*http.Response, error) {
	end := time.Now().Unix()
	elapsed := time.Since(start) / 1000000
	data := map[string]interface{}{
		"start":   int64(start.Unix()),
		"end":     int64(end),
		"elapsed": elapsed,
		"message": message,
		"success": success,
	}
	json_data, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", logUrl, bytes.NewBuffer(json_data))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)
	if err != nil {
		return nil, err
	}
	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)

	return resp, nil
}
