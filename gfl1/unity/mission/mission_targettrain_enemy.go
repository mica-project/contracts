package mission

import "encoding/json/v2"

type MissionTargettrainEnemy struct {
	Code           string `json:"code"`
	Des            string `json:"des"`
	EnemyTeamId    int64  `json:"enemy_team_id"`
	Id             int64  `json:"id"`
	LogFitterName  string `json:"log_fitter_name"`
	Name           string `json:"name"`
	Power          int64  `json:"power"`
	RecommendPower int64  `json:"recommend_power"`
}

type MissionTargettrainEnemyList []MissionTargettrainEnemy

func (s *MissionTargettrainEnemy) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
