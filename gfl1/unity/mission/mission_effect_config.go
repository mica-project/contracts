package mission

import "encoding/json/v2"

type MissionEffectConfig struct {
	Code        string  `json:"code"`
	Description string  `json:"description"`
	Duration    int64   `json:"duration"`
	Id          int64   `json:"id"`
	Name        string  `json:"name"`
	Offset      []int64 `json:"offset"`
	Priority    int64   `json:"priority"`
	Size        float64 `json:"size"`
	Sound       string  `json:"sound"`
}

type MissionEffectConfigList []MissionEffectConfig

func (s *MissionEffectConfig) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
