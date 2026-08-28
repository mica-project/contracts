package squad

import "encoding/json/v2"

type SquadChipExp struct {
	Exp          int64 `json:"exp"`
	Lv           int64 `json:"lv"`
	StrengthCoef int64 `json:"strength_coef"`
}

type SquadChipExpList []SquadChipExp

func (s *SquadChipExp) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
