package breakout

import "encoding/json/v2"

type BreakoutGun struct {
	Armor        string  `json:"armor"`
	Atk          string  `json:"atk"`
	BtTree       string  `json:"bt_tree"`
	Code         string  `json:"code"`
	DeploymentCd int64   `json:"deployment_cd"`
	Dodge        string  `json:"dodge"`
	Hp           string  `json:"hp"`
	Icon         string  `json:"icon"`
	Id           int64   `json:"id"`
	Name         string  `json:"name"`
	Piercing     string  `json:"piercing"`
	Range        string  `json:"range"`
	Rate         string  `json:"rate"`
	Shield       string  `json:"shield"`
	SkillDec     string  `json:"skill_dec"`
	Skills       []int64 `json:"skills"`
	Speed        string  `json:"speed"`
	Type         int64   `json:"type"`
}

type BreakoutGunList []BreakoutGun

func (s *BreakoutGun) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
