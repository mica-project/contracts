package breakout

import "encoding/json/v2"

type BreakoutItem struct {
	CostPart int64    `json:"cost_part"`
	Dec      []string `json:"dec"`
	Icon     string   `json:"icon"`
	Id       int64    `json:"id"`
	Name     string   `json:"name"`
	SkillIds string   `json:"skill_ids"`
	SoldPart int64    `json:"sold_part"`
	Type     int64    `json:"type"`
}

type BreakoutItemList []BreakoutItem

func (s *BreakoutItem) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
