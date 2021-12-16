package utils

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"time"
)

type HttpClient struct {
	baseURL    string
	timeout    int
	httpClient http.Client
}

func (client *HttpClient) GetData(pctx context.Context, query string) ([]byte, error) {
	duration := time.Duration(client.timeout) * time.Second
	req, err := http.NewRequest(http.MethodGet, client.baseURL+query, nil)
	if err != nil {
		LogError("Error creating request: "+err.Error(), "GetData")
		return nil, err
	}
	return client.execute(pctx, duration, req)
}

func (client *HttpClient) execute(pctx context.Context, duration time.Duration, req *http.Request) ([]byte, error) {
	ctx, cancel := context.WithTimeout(pctx, duration)
	defer cancel()
	req = req.WithContext(ctx)
	LogInfo("Executing "+req.Method+" Request. URL: "+req.URL.RequestURI(), "Execute HTTP Request")
	res, err := client.httpClient.Do(req)
	if err != nil {
		LogError("Error executing request: "+err.Error(), "Execute HTTP Request")
		return nil, err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		LogError("Error reading response: "+err.Error(), "Execute HTTP Request")
		return nil, err
	}
	LogInfo("Request Executed", "Execute HTTP Request")
	return data, nil
}

func (client *HttpClient) PostData(pctx context.Context, body []byte) ([]byte, error) {
	duration := time.Duration(client.timeout) * time.Second
	req, err := http.NewRequest(http.MethodPost, client.baseURL, bytes.NewBuffer(body))
	if err != nil {
		LogError("Error creating POST request: "+err.Error(), "PostData")
		return nil, err
	}
	return client.execute(pctx, duration, req)
}

func NewHttpClient(url string, timeout int) *HttpClient {
	return &HttpClient{
		baseURL:    url,
		timeout:    timeout,
		httpClient: http.Client{},
	}
}
