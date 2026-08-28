package mission

import "encoding/json/v2"

type MissionIntensifyItemConfig struct {
	ItemId             int64  `json:"item_id"`
	MissionSkillConfig string `json:"mission_skill_config"`
}

type MissionIntensifyItemConfigList []MissionIntensifyItemConfig

func (s *MissionIntensifyItemConfig) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
