package mission

import "encoding/json/v2"

type MissionSkillConfig struct {
	CdTime         int64  `json:"cd_time"`
	Code           string `json:"code"`
	Consumption    int64  `json:"consumption"`
	DataPool       string `json:"data_pool"`
	Id             int64  `json:"id"`
	Level          int64  `json:"level"`
	SkillGroupId   int64  `json:"skill_group_id"`
	SpecialSpotAdd string `json:"special_spot_add"`
	SpotBelong     string `json:"spot_belong"`
	SpotEchelon    string `json:"spot_echelon"`
	SpotType       string `json:"spot_type"`
}

type MissionSkillConfigList []MissionSkillConfig

func (s *MissionSkillConfig) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
