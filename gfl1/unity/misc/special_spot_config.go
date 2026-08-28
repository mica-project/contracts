package misc

import "encoding/json/v2"

type SpecialSpotConfig struct {
	Code                 string  `json:"code"`
	ConditionsTime       string  `json:"conditions_time"`
	ConditionsType       string  `json:"conditions_type"`
	ConflictType         int64   `json:"conflict_type"`
	Description          string  `json:"description"`
	DetectTargetType     string  `json:"detect_target_type"`
	DurationType         string  `json:"duration_type"`
	EffectConditions     []int64 `json:"effect_conditions"`
	EffectRangeType      int64   `json:"effect_range_type"`
	EffectSpotbelongType []int64 `json:"effect_spotbelong_type"`
	Id                   int64   `json:"id"`
	Name                 string  `json:"name"`
	Priority             int64   `json:"priority"`
	SkillConfigIdOnBirth []int64 `json:"skill_config_id_on_birth"`
	SkillConfigIdOnDeath []int64 `json:"skill_config_id_on_death"`
}

type SpecialSpotConfigList []SpecialSpotConfig

func (s *SpecialSpotConfig) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
