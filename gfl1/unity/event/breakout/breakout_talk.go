package breakout

import "encoding/json/v2"

type BreakoutTalk struct {
	Id      int64  `json:"id"`
	Line    string `json:"line"`
	Trigger int64  `json:"trigger"`
	Weight  int64  `json:"weight"`
}

type BreakoutTalkList []BreakoutTalk

func (s *BreakoutTalk) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
