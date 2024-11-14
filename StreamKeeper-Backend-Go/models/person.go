package models

type Person struct {
	Media
	Name               string        `json:"name"`
	KnownFor           []interface{} `json:"known_for"`
	Gender             int           `json:"gender"`
	KnownForDepartment string        `json:"known_for_department"`
}

func NewPerson(data map[string]interface{}) *Person {
	return &Person{
		Media:              *NewMedia(data),
		Name:               getStringOrDefault(data, "name", ""),
		KnownFor:           data["known_for"].([]interface{}),
		Gender:             int(getFloatOrDefault(data, "gender", 0.0)),
		KnownForDepartment: getStringOrDefault(data, "known_for_department", ""),
	}
}
