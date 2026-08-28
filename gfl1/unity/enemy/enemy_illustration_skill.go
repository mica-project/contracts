package enemy

import "encoding/json/v2"

type EnemyIllustrationSkill struct {
	Description string `json:"description"`
	Id          int64  `json:"id"`
	Name        string `json:"name"`
}

type EnemyIllustrationSkillList []EnemyIllustrationSkill

func (s *EnemyIllustrationSkill) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
