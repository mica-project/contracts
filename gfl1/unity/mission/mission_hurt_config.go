package mission

import "encoding/json/v2"

type MissionHurtConfig struct {
	EffectBirth  int64  `json:"effect_birth"`
	HostageValue int64  `json:"hostage_value"`
	Id           int64  `json:"id"`
	LifeType     int64  `json:"life_type"`
	Name         string `json:"name"`
	TargetType   int64  `json:"target_type"`
	Type         int64  `json:"type"`
	Value        string `json:"value"`
}

type MissionHurtConfigList []MissionHurtConfig

func (s *MissionHurtConfig) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
