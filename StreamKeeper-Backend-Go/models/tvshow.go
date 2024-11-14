package models

type TVShow struct {
	Media
	Name         string  `json:"name"`
	OriginalName string  `json:"original_name"`
	FirstAirDate string  `json:"first_air_date"`
	GenreIDs     []int   `json:"genre_ids"`
	VoteAverage  float64 `json:"vote_average"`
	VoteCount    int     `json:"vote_count"`
}

func NewTVShow(data map[string]interface{}) *TVShow {
	genreIDs := []int{}
	if genres, ok := data["genre_ids"].([]interface{}); ok {
		for _, id := range genres {
			if intID, ok := id.(float64); ok {
				genreIDs = append(genreIDs, int(intID))
			}
		}
	}

	return &TVShow{
		Media:        *NewMedia(data),
		Name:         getStringOrDefault(data, "name", ""),
		OriginalName: getStringOrDefault(data, "original_name", ""),
		FirstAirDate: getStringOrDefault(data, "first_air_date", ""),
		GenreIDs:     genreIDs,
		VoteAverage:  getFloatOrDefault(data, "vote_average", 0.0),
		VoteCount:    int(getFloatOrDefault(data, "vote_count", 0.0)),
	}
}
