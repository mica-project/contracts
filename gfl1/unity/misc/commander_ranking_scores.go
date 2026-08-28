package misc

import "encoding/json/v2"

type CommanderRankingScores struct {
	BasicScores float64 `json:"basic_scores"`
	Code        string  `json:"code"`
	Id          int64   `json:"id"`
	KSlopes     string  `json:"k_slopes"`
	TypeId      int64   `json:"type_id"`
	XCounts     string  `json:"x_counts"`
}

type CommanderRankingScoresList []CommanderRankingScores

func (s *CommanderRankingScores) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
