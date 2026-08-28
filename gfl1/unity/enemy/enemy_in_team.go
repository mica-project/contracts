package enemy

import "encoding/json/v2"

type EnemyInTeam struct {
	CoordinatorX         int64 `json:"coordinator_x"`
	CoordinatorY         int64 `json:"coordinator_y"`
	EnemyCharacterTypeId int64 `json:"enemy_character_type_id"`
	EnemyTeamId          int64 `json:"enemy_team_id"`
	Id                   int64 `json:"id"`
	Level                int64 `json:"level"`
	Number               int64 `json:"number"`
}

type EnemyInTeamList []EnemyInTeam

func (s *EnemyInTeam) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
