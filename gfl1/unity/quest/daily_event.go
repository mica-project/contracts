package quest

import "encoding/json/v2"

type DailyEvent struct {
	Hint    string   `json:"hint"`
	Id      int64    `json:"id"`
	Name    string   `json:"name"`
	Notes   string   `json:"notes"`
	Options []int64  `json:"options"`
	Script  []string `json:"script"`
}

type DailyEventList []DailyEvent

func (s *DailyEvent) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
