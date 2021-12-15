package utils

import "context"

type HttpClient struct {
	baseURL string
}

func (client *HttpClient) GetData(ctx *context.Context, query string) []byte {
	return []byte{}
}

func (client *HttpClient) PostData(ctx *context.Context, body interface{}) []byte {
	return []byte{}
}

func HandleErrorInRequest(err error) {

}

func NewHttpClient(url string) *HttpClient {
	return &HttpClient{
		baseURL: url,
	}
}
