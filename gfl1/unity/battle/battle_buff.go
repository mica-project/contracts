package battle

import "encoding/json/v2"

type BattleBuff struct {
	BossAvailable int64  `json:"boss_available"`
	ConflictType  int64  `json:"conflict_type"`
	Duration      string `json:"duration"`
	DurationType  int64  `json:"duration_type"`
	Id            int64  `json:"id"`
	MaxTier       int64  `json:"max_tier"`
	Name          string `json:"name"`
	Type          int64  `json:"type"`
}

type BattleBuffList []BattleBuff

func (s *BattleBuff) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
