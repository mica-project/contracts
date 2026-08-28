package squad

import "encoding/json/v2"

type SquadType struct {
	AssistDamage   float64 `json:"assist_damage"`
	AssistDefBreak float64 `json:"assist_def_break"`
	AssistHit      float64 `json:"assist_hit"`
	AssistReload   float64 `json:"assist_reload"`
	ClassEnName    string  `json:"class_en_name"`
	ClassName      string  `json:"class_name"`
	EnName         string  `json:"en_name"`
	FixTime        int64   `json:"fix_time"`
	FixType        int64   `json:"fix_type"`
	Hp             int64   `json:"hp"`
	MpFix          int64   `json:"mp_fix"`
	Name           string  `json:"name"`
	PartFix        float64 `json:"part_fix"`
	TypeId         int64   `json:"type_id"`
}

type SquadTypeList []SquadType

func (s *SquadType) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
