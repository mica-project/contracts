package breakout

import "encoding/json/v2"

type BreakoutDrop struct {
	DropItem   string `json:"drop_item"`
	DropTimes  int64  `json:"drop_times"`
	DropWeight string `json:"drop_weight"`
	Id         int64  `json:"id"`
}

type BreakoutDropList []BreakoutDrop

func (s *BreakoutDrop) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
