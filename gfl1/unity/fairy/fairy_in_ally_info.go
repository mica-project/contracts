package fairy

import "encoding/json/v2"

type FairyInAllyInfo struct {
	FairyId      string `json:"fairy_id"`
	FairyLv      string `json:"fairy_lv"`
	Id           string `json:"id"`
	PassiveSkill string `json:"passive_skill"`
	QualityLv    string `json:"quality_lv"`
	SkillLv      string `json:"skill_lv"`
}

type FairyInAllyInfoList []FairyInAllyInfo

func (s *FairyInAllyInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
