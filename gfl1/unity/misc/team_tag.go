package misc

import "encoding/json/v2"

type TeamTag struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type TeamTagList []TeamTag

func (s *TeamTag) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
