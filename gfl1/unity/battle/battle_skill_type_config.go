package battle

import "encoding/json/v2"

type BattleSkillTypeConfig struct {
	ChargeTier      int64 `json:"charge_tier"`
	ChargeTime      int64 `json:"charge_time"`
	Id              int64 `json:"id"`
	StartChargeTier int64 `json:"start_charge_tier"`
}

type BattleSkillTypeConfigList []BattleSkillTypeConfig

func (s *BattleSkillTypeConfig) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
