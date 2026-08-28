package battle

import "encoding/json/v2"

type BattleHurtConfig struct {
	CriticalHitRate string `json:"critical_hit_rate"`
	DamageRatio     string `json:"damage_ratio"`
	Description     string `json:"description"`
	FloatWide       int64  `json:"float_wide"`
	HitType         int64  `json:"hit_type"`
	Id              int64  `json:"id"`
	IsArmor         int64  `json:"is_armor"`
	IsCriticalHit   int64  `json:"is_critical_hit"`
	IsMiss          int64  `json:"is_miss"`
}

type BattleHurtConfigList []BattleHurtConfig

func (s *BattleHurtConfig) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
