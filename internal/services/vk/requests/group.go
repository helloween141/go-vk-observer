package requests

import "net/http"

type GetGroupRequest struct {
	queryParams map[string]string
}

func NewGetGroupRequest(queryParams map[string]string) *GetGroupRequest {
	return &GetGroupRequest{
		queryParams: queryParams,
	}
}

func (req *GetGroupRequest) GetMethod() string {
	return http.MethodGet
}

func (req *GetGroupRequest) GetEndpoint() string {
	return "groups.getById"
}

func (req *GetGroupRequest) GetQueryParams() map[string]string {
	return req.queryParams
}
