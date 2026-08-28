package battle

import "encoding/json/v2"

type BattleTrigger struct {
	Target int64 `json:"target"`
	Type   int64 `json:"type"`
}

type BattleTriggerList []BattleTrigger

func (s *BattleTrigger) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
