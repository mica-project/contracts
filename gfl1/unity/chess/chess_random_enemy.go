package chess

import "encoding/json/v2"

type ChessRandomEnemy struct {
	EnemyId        []int64 `json:"enemy_id"`
	Id             int64   `json:"id"`
	LaunchTimeType int64   `json:"launch_time_type"`
	RandomSpotId   []int64 `json:"random_spot_id"`
	Time           int64   `json:"time"`
}

type ChessRandomEnemyList []ChessRandomEnemy

func (s *ChessRandomEnemy) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
