package responses

type WallResponse struct {
	Response struct {
		Items []struct {
			ID   int64  `json:"id"`
			Text string `json:"text"`
			Date int64  `json:"date"`
		}
		Repost []struct {
			Text string `json:"text"`
		} `json:"copy_history"`
	}
}
