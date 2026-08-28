package breakout

import "encoding/json/v2"

type BreakoutTrigger struct {
	BtTree string `json:"bt_tree"`
	Id     int64  `json:"id"`
}

type BreakoutTriggerList []BreakoutTrigger

func (s *BreakoutTrigger) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
