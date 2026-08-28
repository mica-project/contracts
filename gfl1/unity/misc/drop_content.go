package misc

import "encoding/json/v2"

type DropContent struct {
	ContentType int64 `json:"content_type"`
	Id          int64 `json:"id"`
	PackageId   int64 `json:"package_id"`
	PrizeId     int64 `json:"prize_id"`
}

type DropContentList []DropContent

func (s *DropContent) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
