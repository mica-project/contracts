package chess

import "encoding/json/v2"

type ChessSkillTrigger struct {
	Id        int64  `json:"id"`
	Parameter int64  `json:"parameter"`
	Target    int64  `json:"target"`
	Type      string `json:"type"`
}

type ChessSkillTriggerList []ChessSkillTrigger

func (s *ChessSkillTrigger) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
