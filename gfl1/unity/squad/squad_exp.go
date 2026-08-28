package squad

import "encoding/json/v2"

type SquadExp struct {
	Lv      int64 `json:"lv"`
	Precise int64 `json:"precise"`
}

type SquadExpList []SquadExp

func (s *SquadExp) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
