package fairy

import "encoding/json/v2"

type FairyInAlly struct {
	FairyId      int64 `json:"fairy_id"`
	FairyLv      int64 `json:"fairy_lv"`
	Id           int64 `json:"id"`
	PassiveSkill int64 `json:"passive_skill"`
	QualityLv    int64 `json:"quality_lv"`
	SkillLv      int64 `json:"skill_lv"`
}

type FairyInAllyList []FairyInAlly

func (s *FairyInAlly) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
