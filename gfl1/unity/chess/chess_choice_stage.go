package chess

import "encoding/json/v2"

type ChessChoiceStage struct {
	Cd          int64   `json:"cd"`
	Code        string  `json:"code"`
	Description string  `json:"description"`
	Id          int64   `json:"id"`
	InitTime    []int64 `json:"init_time"`
	Name        string  `json:"name"`
	Parameter   string  `json:"parameter"`
	Probability int64   `json:"probability"`
	Type        int64   `json:"type"`
}

type ChessChoiceStageList []ChessChoiceStage

func (s *ChessChoiceStage) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
