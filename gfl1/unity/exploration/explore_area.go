package exploration

import "encoding/json/v2"

type ExploreArea struct {
	Background []string `json:"background"`
	Id         int64    `json:"id"`
	Name       string   `json:"name"`
}

type ExploreAreaList []ExploreArea

func (s *ExploreArea) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
