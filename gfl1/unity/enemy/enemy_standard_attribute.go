package enemy

import "encoding/json/v2"

type EnemyStandardAttribute struct {
	Armor            float64 `json:"armor"`
	ArmorPiercing    int64   `json:"armor_piercing"`
	DebuffResistance int64   `json:"debuff_resistance"`
	Def              int64   `json:"def"`
	DefBreak         int64   `json:"def_break"`
	Dodge            float64 `json:"dodge"`
	Hit              float64 `json:"hit"`
	Level            int64   `json:"level"`
	Maxlife          float64 `json:"maxlife"`
	Pow              float64 `json:"pow"`
	Shield           int64   `json:"shield"`
	Tenacity         int64   `json:"tenacity"`
}

type EnemyStandardAttributeList []EnemyStandardAttribute

func (s *EnemyStandardAttribute) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
