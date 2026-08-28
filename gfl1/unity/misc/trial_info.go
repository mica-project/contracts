package misc

import "encoding/json/v2"

type TrialInfo struct {
	EnemyLevel  string `json:"enemy_level"`
	EnemyTeamId string `json:"enemy_team_id"`
	EnemyType   string `json:"enemy_type"`
	Id          string `json:"id"`
}

type TrialInfoList []TrialInfo

func (s *TrialInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
