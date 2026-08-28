package quest

import "encoding/json/v2"

type DailyPass struct {
	EndTime        string  `json:"end_time"`
	Gift           []int64 `json:"gift"`
	Gift2          []int64 `json:"gift2"`
	GiftBackground string  `json:"gift_background"`
	Id             int64   `json:"id"`
	Item           string  `json:"item"`
	Pt             int64   `json:"pt"`
	StartTime      []int64 `json:"start_time"`
	Type           int64   `json:"type"`
}

type DailyPassList []DailyPass

func (s *DailyPass) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
