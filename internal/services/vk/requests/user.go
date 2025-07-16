package requests

import "net/http"

type GetUserRequest struct {
	queryParams map[string]string
}

func NewGetUserRequest(queryParams map[string]string) *GetUserRequest {
	return &GetUserRequest{
		queryParams: queryParams,
	}
}

func (req *GetUserRequest) GetMethod() string {
	return http.MethodGet
}

func (req *GetUserRequest) GetEndpoint() string {
	return "users.get"
}

func (req *GetUserRequest) GetQueryParams() map[string]string {
	return req.queryParams
}
