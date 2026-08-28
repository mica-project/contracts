package breakout

import "encoding/json/v2"

type BreakoutEnemy struct {
	Atk      int64  `json:"atk"`
	BtTree   string `json:"bt_tree"`
	Code     string `json:"code"`
	Hp       int64  `json:"hp"`
	Icon     string `json:"icon"`
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Rate     int64  `json:"rate"`
	SkillDec string `json:"skill_dec"`
	Skills   string `json:"skills"`
}

type BreakoutEnemyList []BreakoutEnemy

func (s *BreakoutEnemy) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
