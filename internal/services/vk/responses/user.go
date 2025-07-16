package responses

type UserResponse struct {
	Response []struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}
}
