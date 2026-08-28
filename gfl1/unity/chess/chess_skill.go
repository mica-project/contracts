package chess

import "encoding/json/v2"

type ChessSkill struct {
	CdType              int64  `json:"cd_type"`
	ChipUiControl       int64  `json:"chip_ui_control"`
	Duration            string `json:"duration"`
	Id                  int64  `json:"id"`
	NegativeSkillTarget int64  `json:"negative_skill_target"`
	TargetBuff          string `json:"target_buff"`
	Trigger             string `json:"trigger"`
}

type ChessSkillList []ChessSkill

func (s *ChessSkill) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
