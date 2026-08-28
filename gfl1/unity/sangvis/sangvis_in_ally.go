package sangvis

import "encoding/json/v2"

type SangvisInAlly struct {
	Chip1                  int64 `json:"chip1"`
	Chip2                  int64 `json:"chip2"`
	Favor                  int64 `json:"favor"`
	Id                     int64 `json:"id"`
	Life                   int64 `json:"life"`
	SangvisAdvance         int64 `json:"sangvis_advance"`
	SangvisId              int64 `json:"sangvis_id"`
	SangvisLevel           int64 `json:"sangvis_level"`
	SangvisResolutionLevel int64 `json:"sangvis_resolution_level"`
	SangvisShapeN          int64 `json:"sangvis_shape_n"`
	Skill1                 int64 `json:"skill1"`
	Skill2                 int64 `json:"skill2"`
	Skill3                 int64 `json:"skill3"`
}

type SangvisInAllyList []SangvisInAlly

func (s *SangvisInAlly) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
