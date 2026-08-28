package misc

import "encoding/json/v2"

type ScoreToRateInfo struct {
	ArmorPiercingRate string `json:"armor_piercing_rate"`
	ArmorRate         string `json:"armor_rate"`
	AtkRate           string `json:"atk_rate"`
	DefRate           string `json:"def_rate"`
	DodgeRate         string `json:"dodge_rate"`
	NightViewRate     string `json:"night_view_rate"`
}

type ScoreToRateInfoList []ScoreToRateInfo

func (s *ScoreToRateInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
