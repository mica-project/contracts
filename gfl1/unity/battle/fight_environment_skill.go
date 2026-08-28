package battle

import "encoding/json/v2"

type FightEnvironmentSkill struct {
	Code      string `json:"code"`
	Desc      string `json:"desc"`
	Id        int64  `json:"id"`
	IfDisplay int64  `json:"if_display"`
	Name      string `json:"name"`
}

type FightEnvironmentSkillList []FightEnvironmentSkill

func (s *FightEnvironmentSkill) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
