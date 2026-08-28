package squad

import "encoding/json/v2"

type SquadAdvancedBonus struct {
	AssistDamage int64 `json:"assist_damage"`
	AssistHit    int64 `json:"assist_hit"`
	GroupId      int64 `json:"group_id"`
	Id           int64 `json:"id"`
	Lv           int64 `json:"lv"`
	UnlockNumber int64 `json:"unlock_number"`
}

type SquadAdvancedBonusList []SquadAdvancedBonus

func (s *SquadAdvancedBonus) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
