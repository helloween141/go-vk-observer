package requests

type RequestInterface interface {
	GetMethod() string
	GetEndpoint() string
	GetQueryParams() map[string]string
}
