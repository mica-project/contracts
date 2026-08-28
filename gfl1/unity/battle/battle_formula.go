package battle

import "encoding/json/v2"

type BattleFormula struct {
	Formula string `json:"formula"`
	Id      int64  `json:"id"`
}

type BattleFormulaList []BattleFormula

func (s *BattleFormula) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
