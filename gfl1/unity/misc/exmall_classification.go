package misc

import "encoding/json/v2"

type ExmallClassification struct {
	Code   string `json:"code"`
	EnName string `json:"en_name"`
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Sort   int64  `json:"sort"`
}

type ExmallClassificationList []ExmallClassification

func (s *ExmallClassification) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
