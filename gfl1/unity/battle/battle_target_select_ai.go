package battle

import "encoding/json/v2"

type BattleTargetSelectAi struct {
	BuffTriggerType int64  `json:"buff_trigger_type"`
	Description     string `json:"description"`
	Id              int64  `json:"id"`
	IsIngoreRange   int64  `json:"is_ingore_range"`
	ListRanking     int64  `json:"list_ranking"`
	TargetNumber    int64  `json:"target_number"`
	TargetPriority  int64  `json:"target_priority"`
	TargetType      int64  `json:"target_type"`
}

type BattleTargetSelectAiList []BattleTargetSelectAi

func (s *BattleTargetSelectAi) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
