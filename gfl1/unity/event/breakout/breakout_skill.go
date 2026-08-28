package breakout

import "encoding/json/v2"

type BreakoutSkill struct {
	Action     string `json:"action"`
	Delay      string `json:"delay"`
	Effect     string `json:"effect"`
	Id         int64  `json:"id"`
	Node       string `json:"node"`
	Percentage string `json:"percentage"`
}

type BreakoutSkillList []BreakoutSkill

func (s *BreakoutSkill) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
