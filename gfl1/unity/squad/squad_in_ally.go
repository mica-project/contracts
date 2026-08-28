package squad

import "encoding/json/v2"

type SquadInAlly struct {
	AdvancedLevel        int64 `json:"advanced_level"`
	CpuLevel             int64 `json:"cpu_level"`
	Id                   int64 `json:"id"`
	Rank                 int64 `json:"rank"`
	Skill1               int64 `json:"skill1"`
	Skill2               int64 `json:"skill2"`
	Skill3               int64 `json:"skill3"`
	SquadCompletionLevel int64 `json:"squad_completion_level"`
	SquadId              int64 `json:"squad_id"`
	SquadLevel           int64 `json:"squad_level"`
}

type SquadInAllyList []SquadInAlly

func (s *SquadInAlly) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
