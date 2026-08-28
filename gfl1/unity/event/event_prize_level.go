package event

import "encoding/json/v2"

type EventPrizeLevel struct {
	Background     string `json:"background"`
	Condition      string `json:"condition"`
	EndTime        string `json:"end_time"`
	Id             int64  `json:"id"`
	LeftIcon       string `json:"left_icon"`
	PrizeEndTime   string `json:"prize_end_time"`
	PrizeStartTime string `json:"prize_start_time"`
	StartTime      string `json:"start_time"`
	Type           int64  `json:"type"`
	Value          string `json:"value"`
}

type EventPrizeLevelList []EventPrizeLevel

func (s *EventPrizeLevel) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
