package sangvis

import "encoding/json/v2"

type SangvisExp struct {
	Exp int64 `json:"exp"`
	Lv  int64 `json:"lv"`
}

type SangvisExpList []SangvisExp

func (s *SangvisExp) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
