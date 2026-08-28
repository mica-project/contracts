package misc

import "encoding/json/v2"

type Present struct {
	DuePrizeId int64  `json:"due_prize_id"`
	DueTime    string `json:"due_time"`
	ItemId     int64  `json:"item_id"`
	PrizeId    int64  `json:"prize_id"`
	Type       int64  `json:"type"`
}

type PresentList []Present

func (s *Present) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
