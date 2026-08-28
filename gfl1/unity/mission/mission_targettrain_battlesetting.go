package mission

import "encoding/json/v2"

type MissionTargettrainBattlesetting struct {
	DifficultLevel         int64   `json:"difficult_level"`
	DodgeDefaultLevel      int64   `json:"dodge_default_level"`
	DodgeLevel             []int64 `json:"dodge_level"`
	ForcefieldDefaultLevel int64   `json:"forcefield_default_level"`
	ForcefieldLevel        []int64 `json:"forcefield_level"`
	GuardDefaultLevel      int64   `json:"guard_default_level"`
	GuardLevel             []int64 `json:"guard_level"`
	ShieldDefaultLevel     int64   `json:"shield_default_level"`
	ShieldLevel            []int64 `json:"shield_level"`
}

type MissionTargettrainBattlesettingList []MissionTargettrainBattlesetting

func (s *MissionTargettrainBattlesetting) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
