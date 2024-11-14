package models

type Media struct {
	ID           int     `json:"id"`
	MediaType    string  `json:"media_type"`
	Popularity   float64 `json:"popularity"`
	Overview     string  `json:"overview"`
	PosterPath   string  `json:"poster_path"`
	BackdropPath string  `json:"backdrop_path"`
}

func NewMedia(data map[string]interface{}) *Media {
	return &Media{
		ID:           int(data["id"].(float64)),
		MediaType:    getStringOrDefault(data, "media_type", "unknown"),
		Popularity:   getFloatOrDefault(data, "popularity", 0.0),
		Overview:     getStringOrDefault(data, "overview", ""),
		PosterPath:   getStringOrDefault(data, "poster_path", ""),
		BackdropPath: getStringOrDefault(data, "backdrop_path", ""),
	}
}
