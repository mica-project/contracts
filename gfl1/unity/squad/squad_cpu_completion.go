package squad

import "encoding/json/v2"

type SquadCpuCompletion struct {
	AssistDamage int64 `json:"assist_damage"`
	AssistHit    int64 `json:"assist_hit"`
	GroupId      int64 `json:"group_id"`
	Id           int64 `json:"id"`
	Lv           int64 `json:"lv"`
	UnlockNumber int64 `json:"unlock_number"`
}

type SquadCpuCompletionList []SquadCpuCompletion

func (s *SquadCpuCompletion) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
