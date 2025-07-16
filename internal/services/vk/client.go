package vk

import (
	"encoding/json"
	"go-vk-observer/internal/services/vk/requests"
	"go-vk-observer/internal/services/vk/responses"
	"net/http"
	"time"
)

const timeout = 10
const maxPostsCount = "10"

type Client struct {
	httpClient  *http.Client
	baseUrl     string
	accessToken string
	apiVersion  string
}

func NewClient(baseUrl string, accessToken string, apiVersion string) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: timeout * time.Second,
		},
		baseUrl:     baseUrl,
		accessToken: accessToken,
		apiVersion:  apiVersion,
	}
}

func (client *Client) SendGetWallRequest(slug string) (*responses.WallResponse, error) {
	request := requests.NewGetWallRequest(map[string]string{
		"domain": slug,
		"count":  maxPostsCount,
	})

	resp, err := client.makeRequest(request)
	if err != nil {
		return nil, err
	}

	var responseBody *responses.WallResponse

	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

func (client *Client) SendGetGroupRequest(slug string) (*responses.GroupResponse, error) {
	request := requests.NewGetGroupRequest(map[string]string{
		"group_id": slug,
	})

	resp, err := client.makeRequest(request)
	if err != nil {
		return nil, err
	}

	var responseBody *responses.GroupResponse

	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

func (client *Client) SendGetUserRequest(slug string) (*responses.UserResponse, error) {
	request := requests.NewGetUserRequest(map[string]string{
		"user_ids": slug,
	})

	resp, err := client.makeRequest(request)
	if err != nil {
		return nil, err
	}

	var responseBody *responses.UserResponse

	err = json.NewDecoder(resp.Body).Decode(&responseBody)

	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

func (client *Client) makeRequest(r requests.RequestInterface) (*http.Response, error) {
	req, err := http.NewRequest(r.GetMethod(), client.baseUrl+"/method/"+r.GetEndpoint(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+client.accessToken)

	q := req.URL.Query()
	for key, value := range r.GetQueryParams() {
		q.Add(key, value)
	}
	q.Add("v", client.apiVersion)

	req.URL.RawQuery = q.Encode()

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
