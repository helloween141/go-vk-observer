package responses

type WallResponse struct {
	Response struct {
		Posts []PostInfo `json:"items"`
	} `json:"response"`
}

type PostInfo struct {
	ID         int64  `json:"id"`
	FromID     int64  `json:"from_id"`
	Text       string `json:"text"`
	Date       int64  `json:"date"`
	RepostInfo []struct {
		Text string `json:"text"`
	} `json:"copy_history"`
}
