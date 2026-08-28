package misc

import "encoding/json/v2"

type ExtraSpine struct {
	Ai         int64  `json:"ai"`
	ExploreTag string `json:"explore_tag"`
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Scale      int64  `json:"scale"`
	SpineCode  string `json:"spine_code"`
}

type ExtraSpineList []ExtraSpine

func (s *ExtraSpine) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
