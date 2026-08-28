package enemy

import "encoding/json/v2"

type EnemyTeam struct {
	EnemyLeader           int64 `json:"enemy_leader"`
	Id                    int64 `json:"id"`
	TargettrainCancollect int64 `json:"targettrain_cancollect"`
}

type EnemyTeamList []EnemyTeam

func (s *EnemyTeam) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
