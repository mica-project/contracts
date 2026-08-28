package enemy

import "encoding/json/v2"

type EnemyCharacterType struct {
	Character       string `json:"character"`
	DeploymentScale int64  `json:"deployment_scale"`
	EffectRatio     int64  `json:"effect_ratio"`
	Id              int64  `json:"id"`
	Level           int64  `json:"level"`
	NormalAttack    int64  `json:"normal_attack"`
	PassiveSkill    string `json:"passive_skill"`
}

type EnemyCharacterTypeList []EnemyCharacterType

func (s *EnemyCharacterType) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
