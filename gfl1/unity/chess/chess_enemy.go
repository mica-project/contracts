package chess

import "encoding/json/v2"

type ChessEnemy struct {
	Ap             int64    `json:"ap"`
	AttackAngle    int64    `json:"attack_angle"`
	AttackTimes    int64    `json:"attack_times"`
	Code           string   `json:"code"`
	Hp             int64    `json:"hp"`
	Id             int64    `json:"id"`
	LifebarOffset  []int64  `json:"lifebar_offset"`
	Name           string   `json:"name"`
	Power          int64    `json:"power"`
	Range          int64    `json:"range"`
	RewardChipIds  []string `json:"reward_chip_ids"`
	SpineDirection int64    `json:"spine_direction"`
}

type ChessEnemyList []ChessEnemy

func (s *ChessEnemy) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
