package misc

import "encoding/json/v2"

type TriggerIndex struct {
	Id   int64 `json:"id"`
	Type int64 `json:"type"`
}

type TriggerIndexList []TriggerIndex

func (s *TriggerIndex) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
