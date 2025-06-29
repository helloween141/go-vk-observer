package requests

import "net/http"

type GetWallRequest struct {
	queryParams map[string]string
}

func NewGetWallRequest(queryParams map[string]string) *GetWallRequest {
	return &GetWallRequest{
		queryParams: queryParams,
	}
}

func (req *GetWallRequest) GetMethod() string {
	return http.MethodGet
}

func (req *GetWallRequest) GetEndpoint() string {
	return "wall.get"
}

func (req *GetWallRequest) GetQueryParams() map[string]string {
	return req.queryParams
}
