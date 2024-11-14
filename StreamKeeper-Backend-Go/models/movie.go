package models

type Movie struct {
	Media
	Title         string  `json:"title"`
	OriginalTitle string  `json:"original_title"`
	ReleaseDate   string  `json:"release_date"`
	GenreIDs      []int   `json:"genre_ids"`
	VoteAverage   float64 `json:"vote_average"`
	VoteCount     int     `json:"vote_count"`
}

func NewMovie(data map[string]interface{}) *Movie {
	genreIDs := []int{}
	if genres, ok := data["genre_ids"].([]interface{}); ok {
		for _, id := range genres {
			if intID, ok := id.(float64); ok {
				genreIDs = append(genreIDs, int(intID))
			}
		}
	}

	return &Movie{
		Media:         *NewMedia(data),
		Title:         getStringOrDefault(data, "title", ""),
		OriginalTitle: getStringOrDefault(data, "original_title", ""),
		ReleaseDate:   getStringOrDefault(data, "release_date", ""),
		GenreIDs:      genreIDs,
		VoteAverage:   getFloatOrDefault(data, "vote_average", 0.0),
		VoteCount:     int(getFloatOrDefault(data, "vote_count", 0.0)),
	}
}
