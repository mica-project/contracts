package misc

import "encoding/json/v2"

type RecommendedFormula struct {
	Ammo        int64    `json:"ammo"`
	Background  string   `json:"background"`
	DevelopType int64    `json:"develop_type"`
	Id          int64    `json:"id"`
	Mp          int64    `json:"mp"`
	Mre         int64    `json:"mre"`
	Name        string   `json:"name"`
	Part        int64    `json:"part"`
	Preview     []string `json:"preview"`
	Type        string   `json:"type"`
	TypeRarity  []string `json:"type_rarity"`
}

type RecommendedFormulaList []RecommendedFormula

func (s *RecommendedFormula) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
