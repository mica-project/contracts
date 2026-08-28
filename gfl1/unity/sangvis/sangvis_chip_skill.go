package sangvis

import "encoding/json/v2"

type SangvisChipSkill struct {
	BattleSkill string  `json:"battle_skill"`
	Code        string  `json:"code"`
	Id          int64   `json:"id"`
	TargetRange int64   `json:"target_range"`
	TargetType  []int64 `json:"target_type"`
}

type SangvisChipSkillList []SangvisChipSkill

func (s *SangvisChipSkill) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
