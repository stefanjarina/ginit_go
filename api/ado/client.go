package ado

import (
    "net/http"
    "time"
)

type AzureDevOps struct {
	baseUrl    string
	httpClient *http.Client
	list       []string
}

func NewClient() *AzureDevOps {
	client := &http.Client{
		Timeout: 5 * time.Second,
		}
		return &AzureDevOps{
			baseUrl:    "https://www.toptal.com/developers/gitignore/api",
			httpClient: client,
			}
}