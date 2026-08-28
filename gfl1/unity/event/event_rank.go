package event

import "encoding/json/v2"

type EventRank struct {
	Id     int64  `json:"id"`
	Title  string `json:"title"`
	Weight int64  `json:"weight"`
}

type EventRankList []EventRank

func (s *EventRank) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
