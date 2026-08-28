package chess

import "encoding/json/v2"

type ChessGunType struct {
	Ap           int64  `json:"ap"`
	AttackAngle  int64  `json:"attack_angle"`
	AttackTimes  int64  `json:"attack_times"`
	GunTypeSkill string `json:"gun_type_skill"`
	Hp           int64  `json:"hp"`
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Power        int64  `json:"power"`
	Range        int64  `json:"range"`
}

type ChessGunTypeList []ChessGunType

func (s *ChessGunType) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
