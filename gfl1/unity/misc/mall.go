package misc

import "encoding/json/v2"

type Mall struct {
	ClassificationId int64 `json:"classification_id"`
	Discount         int64 `json:"discount"`
	Gemprice         int64 `json:"gemprice"`
	Id               int64 `json:"id"`
	Mp               int64 `json:"mp"`
	Type             int64 `json:"type"`
}

type MallList []Mall

func (s *Mall) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
