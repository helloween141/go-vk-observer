package responses

type GroupResponse struct {
	Response struct {
		Groups []struct {
			Name string `json:"name"`
		}
	}
}
