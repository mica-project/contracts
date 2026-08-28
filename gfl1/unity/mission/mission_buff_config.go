package mission

import "encoding/json/v2"

type MissionBuffConfig struct {
	Code           string `json:"code"`
	ConflictType   int64  `json:"conflict_type"`
	Description    string `json:"description"`
	DurationTime   string `json:"duration_time"`
	DurationType   string `json:"duration_type"`
	ExpBuff        string `json:"exp_buff"`
	Id             int64  `json:"id"`
	IsResourceBuff int64  `json:"is_resource_buff"`
	Name           string `json:"name"`
	Type           int64  `json:"type"`
}

type MissionBuffConfigList []MissionBuffConfig

func (s *MissionBuffConfig) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
