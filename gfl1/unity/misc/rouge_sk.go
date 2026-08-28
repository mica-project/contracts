package misc

import "encoding/json/v2"

type RougeSk struct {
	Code    string `json:"code"`
	Id      int64  `json:"id"`
	RankC   int64  `json:"rank_c"`
	SkillId int64  `json:"skill_id"`
	TypeC   string `json:"type_c"`
}

type RougeSkList []RougeSk

func (s *RougeSk) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
