package chess

import "encoding/json/v2"

type ChessChipTargetSelect struct {
	Id           int64  `json:"id"`
	SelectOrder  string `json:"select_order"`
	SelectType   int64  `json:"select_type"`
	TargetNumber int64  `json:"target_number"`
	TargetType   string `json:"target_type"`
}

type ChessChipTargetSelectList []ChessChipTargetSelect

func (s *ChessChipTargetSelect) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
