package models

type Review struct {
	Author  string  `json:"author"`
	Content string  `json:"content"`
	Created string  `json:"created_at"`
	Updated string  `json:"updated_at"`
	Rating  float64 `json:"rating"`
}

func NewReview(data map[string]interface{}) *Review {
	authorDetails := data["author_details"].(map[string]interface{})
	rating := 0.0
	if r, ok := authorDetails["rating"].(float64); ok {
		rating = r
	}

	return &Review{
		Author:  getStringOrDefault(data, "author", ""),
		Content: getStringOrDefault(data, "content", ""),
		Created: getStringOrDefault(data, "created_at", ""),
		Updated: getStringOrDefault(data, "updated_at", ""),
		Rating:  rating,
	}
}
